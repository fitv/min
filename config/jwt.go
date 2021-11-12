package config

import (
	"time"

	"github.com/fitv/min/core/config"
)

type JwtConfig struct {
	SigningKey string
	Secret     string
	TTL        time.Duration
}

var Jwt = &JwtConfig{
	SigningKey: "jwtClaims",
	TTL:        time.Hour * 1,
	Secret:     config.GetString("jwt.secret"),
}
