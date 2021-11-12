package web

import (
	"net/http"

	"github.com/fitv/min/config"
	"github.com/gin-gonic/gin"
)

type Home struct{}

func (Home) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": config.App.Name,
	})
}
