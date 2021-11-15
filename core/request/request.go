package request

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Token returns the token from the request header or query.
func Token(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	splits := strings.SplitN(strings.TrimSpace(token), " ", 2)
	if len(splits) == 2 {
		return splits[1]
	}
	return c.Query("token")
}

// IsApiRoute determines whether the request path is an api route.
func IsApiRoute(c *gin.Context) bool {
	return strings.HasPrefix(c.Request.URL.Path, "/api/")
}
