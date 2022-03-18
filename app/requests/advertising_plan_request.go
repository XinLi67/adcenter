package requests

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdvertisingPlanRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	CreatorId             uint64    `valid:"creator_id" json:"creator_id,omitempty"`
	AdvertisingId         uint64    `valid:"advertising_id" json:"advertising_id,omitempty"`
	AdvertisingType       uint64    `valid:"advertising_type" json:"advertising_type,omitempty"`
	AdvertisingPositionId uint64    `valid:"advertising_position_id" json:"advertising_position_id,omitempty"`
	Order                 uint64    `valid:"order" json:"order,omitempty"`
	SchedulingDate        uint64    `valid:"scheduling_date" json:"scheduling_date,omitempty"`
	SchedulingTime        uint64    `valid:"scheduling_time" json:"scheduling_time,omitempty"`
	StartDate             time.Time `valid:"start_date" json:"start_date,omitempty"`
	EndTDate              time.Time `valid:"end_date" json:"end_date,omitempty"`
	StartTime             time.Time `valid:"start_time" json:"start_time,omitempty"`
	EndTime               time.Time `valid:"end_time" json:"end_time,omitempty"`
	Status                uint64    `valid:"status" json:"status,omitempty"`
}

func AdvertisingPlanSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:advertising_plans,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"advertising_type": []string{"numeric_between:-1,4"},
		"SchedulingDate":   []string{"numeric_between:-1,2"},
		"SchedulingTime":   []string{"numeric_between:-1,2"},
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
		"advertising_type": []string{
			"numeric_between:只能为0或1或2或3",
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
