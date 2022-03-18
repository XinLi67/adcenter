package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type MediaRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	CreatorId    uint64 `valid:"creator_id" json:"creator_id,omitempty"`
	MediaGroupId uint64 `valid:"media_group" json:"media_group,omitempty"`
	DepartmentId uint64 `valid:"department_id" json:"department_id,omitempty"`
	Type         uint64 `valid:"type" json:"type,omitempty"`
	Url          string `valid:"url" json:"url,omitempty"`
	Title        string `valid:"title" json:"title,omitempty"`
	Content      string `valid:"content" json:"content,omitempty"`
}

func MediaSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:media,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"type":  []string{"numeric_between:-1,2"},
		"url":   []string{"url"},
		"title": []string{"min:2", "max:60"},
	}
	messages := govalidator.MapData{
		// "name": []string{
		//     "required:名称为必填项",
		//     "min_cn:名称长度需至少 2 个字",
		//     "max_cn:名称长度不能超过 8 个字",
		//     "not_exists:名称已存在",
		// },
		// "description": []string{
		//     "min_cn:描述长度需至少 3 个字",
		//     "max_cn:描述长度不能超过 255 个字",
		// },
		"type": []string{
			"numeric_between:只能为0或1",
		},
		"url": []string{
			"url:必须是有效的URL",
		},
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
		},
	}
	return validate(data, rules, messages)
}
