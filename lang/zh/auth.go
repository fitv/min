package zh

import "github.com/fitv/min/core/lang/zh"

func init() {
	zh.Set("auth", map[string]interface{}{
		"invalid_credentials": "用户名或密码错误",
	})
}
