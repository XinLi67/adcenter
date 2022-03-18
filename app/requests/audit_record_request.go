package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AuditRecordRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	AuditableId   uint64 `valid:"auditable_id" json:"auditable_id,omitempty"`
	AuditableType uint64 `valid:"auditable_type" json:"auditable_type,omitempty"`
	ApplicantId   uint64 `valid:"applicant_id" json:"applicant_id,omitempty"`
	AuditorId     uint64 `valid:"auditor_id" json:"auditor_id,omitempty"`
	Status        uint64 `valid:"status" json:"status,omitempty"`
	Content       string `valid:"content" json:"content,omitempty"`
}

func AuditRecordSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:audit_records,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"auditable_type": []string{"numeric_between:-1,5"},
		"Status":         []string{"numeric_between:-1,2"},
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
		"auditable_type": []string{
			"numeric_between:只能为0或1或2或3或4",
		},
		"Status": []string{
			"numeric_between:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
