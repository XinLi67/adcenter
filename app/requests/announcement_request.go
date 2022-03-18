package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AnnouncementRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	AnnouncementNo uint64 `valid:"announcement_no" json:"advertising_no,omitempty"`
	CreatorId      uint64 `valid:"creator_id" json:"creator_id,omitempty"`
	DepartmentId   uint64 `valid:"department_id" json:"department_id,omitempty"`
	Title          string `valid:"title" json:"title,omitempty"`
	LongTitle      string `valid:"long_title" json:"long_title,omitempty"`
	Type           uint64 `valid:"type" json:"type,omitempty"`
	Banner         string `valid:"banner" json:"banner,omitempty"`
	RedirectTo     uint64 `valid:"redirect_to" json:"redirect_to,omitempty"`
	RedirectParams string `valid:"redirect_params" json:"redirect_params,omitempty"`
	Content        string `valid:"content" json:"content,omitempty"`
}

func AnnouncementSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:announcements,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"title":       []string{"min:2", "max:60"},
		"long_title":  []string{"max:255"},
		"type":        []string{"numeric_between:-1,3"},
		"banner":      []string{"max:255"},
		"redirect_to": []string{"numeric_between:-1,2"},
		"content":     []string{"max:255"},
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
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
		},
		"long_title": []string{
			"max:最大长度为255",
		},
		"type": []string{
			"numeric_between:只能为0或1或2",
		},
		"banner": []string{
			"max:最大长度为255",
		},
		"redirect_to": []string{
			"numeric_between:只能为0或1",
		},
		"content": []string{
			"max:最大长度为255",
		},
	}
	return validate(data, rules, messages)
}
