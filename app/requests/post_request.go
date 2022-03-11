package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PostRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()

	Title   string `valid:"title" json:"title"`
	Content string `valid:"content" json:"content,omitempty"`
}

func PostSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":   []string{"required", "min_cn:2", "max_cn:8", "not_exists:posts,title"},
		"content": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
			"not_exists:名称已存在",
		},
		"content": []string{
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
