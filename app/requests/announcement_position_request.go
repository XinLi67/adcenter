package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AnnouncementPositionRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	AnnouncementPositionName string `valid:"announcement_position_name" json:"announcement_position_name,omitempty"`
	ParentId                 uint64 `valid:"parent_id" json:"parent_id,omitempty"`
	Code                     string `valid:"code" json:"code,omitempty"`
	Height                   uint64 `valid:"height" json:"height,omitempty"`
	Weight                   uint64 `valid:"weight" json:"weight,omitempty"`
}

func AnnouncementPositionSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:announcement_positions,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"height": []string{"numeric"},
		"weight": []string{"numeric"},
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
		"height:": []string{
			"numeric:必须是数字",
		},
		"weight:": []string{
			"numeric:必须是数字",
		},
	}
	return validate(data, rules, messages)
}
