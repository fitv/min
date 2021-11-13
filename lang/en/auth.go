package en

import "github.com/fitv/min/core/lang/en"

func init() {
	en.Set("auth", map[string]interface{}{
		"unauthorized":        "unauthorized",
		"invalid_credentials": "invalid credentials",
		"username_existed":    "username already exists",
	})
}
