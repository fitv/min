package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fitv/min/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// CheckToken check the Request token
func Check(c *gin.Context) (*jwt.RegisteredClaims, bool) {
	if claims, exist := c.Get(config.Jwt.SigningKey); exist {
		if claims, ok := claims.(*jwt.RegisteredClaims); ok {
			return claims, true
		}
	}
	return nil, false
}

// MustUID get the user ID from current request token, panic if unauthorized
func MustUID(c *gin.Context) int {
	claims, ok := Check(c)
	if !ok {
		panic("unauthorized")
	}
	return cast.ToInt(claims.Subject)
}

// SignToken generate JWT token
func SignToken(uid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ID:        uuid.NewString(),
		Subject:   strconv.Itoa(uid),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Jwt.TTL)),
	})
	return token.SignedString([]byte(config.Jwt.Secret))
}

// VerifyToken verify JWT token
func VerifyToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
