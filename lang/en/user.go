package en

import "github.com/fitv/min/core/lang/en"

func init() {
	en.Set("user", map[string]interface{}{
		"not_found": "user not exists",
		"username":  "username",
		"password":  "password",
	})
}
