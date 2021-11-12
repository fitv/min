package routes

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/fitv/min/app/middleware"
	"github.com/fitv/min/app/web"
	"github.com/fitv/min/global"
	"github.com/gin-gonic/gin"
)

func Web(c *gin.Engine) {
	staticFS, _ := fs.Sub(global.FS(), "static")
	favicon, _ := global.FS().ReadFile("static/favicon.ico")
	c.SetHTMLTemplate(template.Must(template.ParseFS(global.FS(), "templates/*")))

	home := web.Home{}
	apidoc := web.ApiDoc{}

	c.GET("/", home.Index)

	c.GET("/apidoc", apidoc.Index)

	static := c.Group("/", middleware.Static())
	{
		static.StaticFS("/static", http.FS(staticFS))

		static.GET("/favicon.ico", func(c *gin.Context) {
			c.Data(http.StatusOK, "image/x-icon", favicon)
		})
	}
}
