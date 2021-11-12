package middleware

import (
	"net/http"
	"strings"

	"github.com/fitv/min/config"
	"github.com/fitv/min/core/auth"
	"github.com/fitv/min/core/lang"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			token = c.Query("token")
		}

		claims, err := auth.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": lang.Trans("auth.unauthorized"),
			})
			return
		}
		c.Set(config.Jwt.SigningKey, claims)

		c.Next()
	}
}
