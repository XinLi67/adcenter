package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdvertisingRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	//FIXME()
	AdvertisingNo  uint64 `valid:"advertising_no" json:"advertising_no,omitempty"`
	CreatorId      uint64 `valid:"creator_id" json:"creator_id,omitempty"`
	DepartmentId   uint64 `valid:"department_id" json:"department_id,omitempty"`
	Title          string `valid:"title" json:"title,omitempty"`
	Type           uint64 `valid:"type" json:"type,omitempty"`
	RedirectTo     uint64 `valid:"redirect_to" json:"redirect_to,omitempty"`
	MediaId        uint64 `valid:"media_id" json:"media_id,omitempty"`
	MediaType      uint64 `valid:"media_type" json:"media_type,omitempty"`
	Size           string `valid:"size" json:"size,omitempty"`
	RedirectParams string `valid:"redirect_params" json:"redirect_params,omitempty"`
	Description    string `valid:"description" json:"description,omitempty"`
}

func AdvertisingSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:advertisings,name"},
		// "description": []string{"min_cn:3", "max_cn:255"},
		"title":       []string{"min:2", "max:60"},
		"type":        []string{"numeric_between:-1,3"},
		"redirect_to": []string{"numeric_between:-1,2"},
		"media_type":  []string{"numeric_between:-1,2"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
		},
		"type": []string{
			"numeric_between:只能为0或1或2",
		},
		"redirect_to": []string{
			"numeric_between:只能为0或1",
		},
		"media_type": []string{
			"numeric_between:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
