package requests

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AnnouncementPlanRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	CreatorId              uint64    `valid:"creator_id" json:"creator_id,omitempty"`
	AnnouncementId         uint64    `valid:"announcement_id" json:"announcement_id,omitempty"`
	AnnouncementType       uint64    `valid:"announcement_type" json:"announcement_type,omitempty"`
	AnnouncementPositionId uint64    `valid:"announcement_position_id" json:"announcement_position_id,omitempty"`
	Order                  uint64    `valid:"order" json:"order,omitempty"`
	SchedulingDate         uint64    `valid:"scheduling_date" json:"scheduling_date,omitempty"`
	SchedulingTime         uint64    `valid:"scheduling_time" json:"scheduling_time,omitempty"`
	StartDate              time.Time `valid:"start_date" json:"start_date,omitempty"`
	EndTDate               time.Time `valid:"end_date" json:"end_date,omitempty"`
	StartTime              time.Time `valid:"start_time" json:"start_time,omitempty"`
	EndTime                time.Time `valid:"end_time" json:"end_time,omitempty"`
	Status                 uint64    `valid:"status" json:"status,omitempty"`
}

func AnnouncementPlanSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:announcement_plans,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"announcement_type": []string{"numeric_between:-1,3"},
		"SchedulingDate":    []string{"numeric_between:-1,2"},
		"SchedulingTime":    []string{"numeric_between:-1,2"},
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
		"announcement_type": []string{
			"numeric_between:只能为0或1或2",
		},
		"SchedulingDate": []string{
			"numeric_between:只能为0或1",
		},
		"SchedulingTime": []string{
			"numeric_between:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
