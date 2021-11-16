package config

import (
	"time"

	"github.com/fitv/min/core/config"
)

type JwtConfig struct {
	SigningKey string // The name of the key used to sign the token in the request context
	Secret     string
	TTL        time.Duration
}

var Jwt = &JwtConfig{
	SigningKey: "jwtClaims",
	Secret:     config.GetString("jwt.secret"),
	TTL:        time.Hour * 1,
}
