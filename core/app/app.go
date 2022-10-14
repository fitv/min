package app

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fitv/go-i18n"
	"github.com/fitv/go-logger"
	"github.com/fitv/min/config"
	"github.com/fitv/min/core/cache"
	"github.com/fitv/min/core/db"
	"github.com/fitv/min/core/redis"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	Version = "1.0.0"
)

type Application struct {
	Gin           *gin.Engine
	FS            embed.FS
	DB            *db.DB
	Cache         cache.Cache
	Redis         *redis.Redis
	Logger        *logger.Logger
	Translator    ut.Translator
	Lang          *i18n.I18n
	Services      []Service
	Routes        []Route
	Validations   []Validation
	ShutdownFuncs []func()
	Version       string
}

type Route func(*gin.Engine)

type Validation func(*validator.Validate, ut.Translator)

type Service interface {
	Boot(*Application)
	Register(*Application)
}

// NewApplication create a new application
func NewApplication(fs embed.FS) *Application {
	return &Application{
		FS:      fs,
		Version: Version,
	}
}

// Run start application
func (app *Application) Run() {
	app.registerService()
	app.bootService()
	app.listenAndServe()
}

// AddShutdown add shutdown function
func (app *Application) AddShutdown(fn func()) {
	app.ShutdownFuncs = append(app.ShutdownFuncs, fn)
}

// AddValidation add validation
func (app *Application) AddValidation(vs ...Validation) {
	app.Validations = append(app.Validations, vs...)
}

// AddService add service
func (app *Application) AddService(ss ...Service) {
	app.Services = append(app.Services, ss...)
}

// Shutdown shutdown the application
func (app *Application) Shutdown() {
	for _, ShutdownFunc := range app.ShutdownFuncs {
		ShutdownFunc()
	}
}

// registerService register the application services
func (app *Application) registerService() {
	for _, service := range app.Services {
		service.Register(app)
	}
}

// bootService boot the application services
func (app *Application) bootService() {
	for _, service := range app.Services {
		service.Boot(app)
	}
}

// listenAndServe start http server
func (app *Application) listenAndServe() {
	srv := &http.Server{
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		Addr:           fmt.Sprintf("%s:%d", config.App.Addr, config.App.Port),
		Handler:        app.Gin,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()
	log.Printf("Server started, Listen %s\n", srv.Addr)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
