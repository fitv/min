package routes

import (
	v1 "github.com/fitv/min/app/api/v1"
	"github.com/fitv/min/app/middleware"
	"github.com/gin-gonic/gin"
)

func Api(c *gin.Engine) {
	auth := &v1.Auth{}
	migrate := &v1.Migrate{}

	v1 := c.Group("/api/v1", middleware.Cors())
	{
		v1.GET("/migrate", migrate.Index)

		v1.POST("/auth/login", auth.Login)
		v1.POST("/auth/register", auth.Register)

		authorized := v1.Group("/", middleware.Auth())
		{
			authorized.GET("/auth/profile", auth.Profile)
			authorized.POST("/auth/refresh", auth.Refresh)
		}
	}
}
