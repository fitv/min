package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/lang"
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
		app.AddClose(func() {
			file.Close()
		})

		gin.DefaultWriter = io.MultiWriter(file)
		gin.DefaultErrorWriter = io.MultiWriter(file)
	}

	app.Gin = gin.New()

	// Register Logger and Recovery middleware
	app.Gin.Use(gin.Logger(), gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": lang.Trans("message.server_error"),
			})
			return
		}

		c.HTML(http.StatusInternalServerError, "500.tmpl", gin.H{})
		c.Abort()
	}))
}
