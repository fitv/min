package zh

import "github.com/fitv/min/core/lang"

func init() {
	lang.Set("auth", map[string]interface{}{
		"unauthorized":        "unauthorized",
		"invalid_credentials": "invalid credentials",
		"username_existed":    "username already exists",
	})
}
