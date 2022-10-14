package service

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/request"
	"github.com/fitv/min/core/response"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Service
}

func (Gin) Register(app *app.Application) {
	if !config.App.Debug {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)

		file, err := os.OpenFile(config.Log.Path+"/gin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			panic(fmt.Errorf("open log file error: %w", err))
		}
		app.AddShutdown(func() {
			file.Close()
		})

		gin.DefaultWriter = io.MultiWriter(file)
		gin.DefaultErrorWriter = io.MultiWriter(file)
	}

	app.Gin = gin.New()

	// Register Logger and Recovery middleware
	app.Gin.Use(gin.Logger(), gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		if request.IsApiRoute(c) {
			response.ServerError(c)
			return
		}
		c.HTML(http.StatusInternalServerError, "500.tmpl", gin.H{})
		c.Abort()
	}))
}
