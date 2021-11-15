package service

import (
	"net/http"

	"github.com/fitv/min/core/app"
	"github.com/fitv/min/core/lang"
	"github.com/fitv/min/core/request"
	"github.com/fitv/min/core/response"
	"github.com/fitv/min/routes"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Service
}

func (Router) Boot(app *app.Application) {
	routes.Web(app.Gin)

	routes.Api(app.Gin)

	// Register the Not Found handler.
	app.Gin.NoRoute(func(c *gin.Context) {
		if request.IsApiRoute(c) {
			response.NotFound(c, lang.Trans("message.not_found"))
			return
		}
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{})
		c.Abort()
	})
}
