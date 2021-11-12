package middleware

import (
	"github.com/gin-gonic/gin"
)

func Static() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")

		c.Next()
	}
}
