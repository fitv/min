package middleware

import (
	"github.com/fitv/min/config"
	"github.com/fitv/min/core/auth"
	"github.com/fitv/min/core/request"
	"github.com/fitv/min/core/response"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := auth.VerifyToken(request.Token(c))
		if err != nil {
			response.Unauthorized(c)
			return
		}
		c.Set(config.Jwt.SigningKey, claims)

		c.Next()
	}
}
