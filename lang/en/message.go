package en

import "github.com/fitv/min/core/lang/en"

func init() {
	en.Set("message", map[string]interface{}{
		"success":         "success",
		"failed":          "failed",
		"not_found":       "not found",
		"forbidden":       "forbidden",
		"server_error":    "server error",
		"validate_failed": "verification failed",
	})
}
