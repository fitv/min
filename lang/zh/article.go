package zh

import "github.com/fitv/min/core/lang"

func init() {
	lang.Set("article", map[string]interface{}{
		"not_found": "文章不存在",
		"title":     "标题",
		"content":   "内容",
		"author":    "作者",
	})
}
