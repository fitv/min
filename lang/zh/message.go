package zh

import "github.com/fitv/min/core/lang/zh"

func init() {
	zh.Set("message", map[string]interface{}{
		"success":         "操作成功",
		"failed":          "操作失败",
		"not_found":       "未找到",
		"unauthorized":    "未登录",
		"forbidden":       "无权限",
		"server_error":    "服务器错误",
		"validate_failed": "验证失败",
	})
}
