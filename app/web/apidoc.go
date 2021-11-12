package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiDoc struct{}

func (ApiDoc) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "apidoc.tmpl", gin.H{
		"version": time.Now().Unix(),
	})
}
