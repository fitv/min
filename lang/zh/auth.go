package zh

import "github.com/fitv/min/core/lang/zh"

func init() {
	zh.Set("auth", map[string]interface{}{
		"unauthorized":        "未登录",
		"invalid_credentials": "用户名或密码错误",
		"username_existed":    "用户名已存在",
	})
}
