package zh

import "github.com/fitv/min/core/lang/zh"

func init() {
	zh.Set("user", map[string]interface{}{
		"not_found": "用户不存在",
		"username":  "姓名",
		"password":  "密码",
	})
}
