package zh

import "github.com/fitv/min/core/lang"

func init() {
	lang.Set("user", map[string]interface{}{
		"not_found": "user not exists",
		"username":  "username",
		"password":  "password",
	})
}
