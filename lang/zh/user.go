package zh

import "github.com/fitv/min/core/lang"

func init() {
	lang.Set("user", map[string]interface{}{
		"not_found": "用户不存在",
		"username":  "姓名",
		"password":  "密码",
	})
}
