package v1

import (
	"github.com/fitv/min/config"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewAccessToken(token string) *AccessToken {
	return &AccessToken{
		AccessToken: token,
		TokenType:   "bearer",
		ExpiresIn:   int(config.Jwt.TTL.Seconds()),
	}
}
