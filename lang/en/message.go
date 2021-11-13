package zh

import "github.com/fitv/min/core/lang"

func init() {
	lang.Set("message", map[string]interface{}{
		"success":         "success",
		"failed":          "failed",
		"not_found":       "not found",
		"forbidden":       "forbidden",
		"server_error":    "server error",
		"validate_failed": "verification failed",
	})
}
