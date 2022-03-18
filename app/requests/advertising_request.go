package requests

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/thedevsaddam/govalidator"
// )

// type AdvertisingRequest struct {
// 	// Name        string `valid:"name" json:"name"`
// 	// Description string `valid:"description" json:"description,omitempty"`
// 	//FIXME()
// 	CreatorId      int    `valid:"creator_id" json:"creator_id,omitempty"`
// 	DepartmentId   int    `valid:"department_id" json:"department_id,omitempty"`
// 	Title          string `valid:"title" json:"title,omitempty"`
// 	Type           int    `valid:"type" json:"type,omitempty"`
// 	RedirectTo     int    `valid:"redirect_to" json:"redirect_to,omitempty"`
// 	MediaId        int    `valid:"media_id" json:"media_id,omitempty"`
// 	MediaType      int    `valid:"media_type" json:"media_type,omitempty"`
// 	Size           string `valid:"size" json:"size,omitempty"`
// 	RedirectParams string `valid:"redirect_params" json:"redirect_params,omitempty"`
// 	Description    string `valid:"description" json:"description,omitempty"`
// }

// func AdvertisingSave(data interface{}, c *gin.Context) map[string][]string {

// 	rules := govalidator.MapData{
// 		// "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:advertisings,name"},
// 		// "description": []string{"min_cn:3", "max_cn:255"},
//         "CreatorId":       int    `gorm:"type:int;column:creator_id;not null"`,
// 		"DepartmentId":   int    `gorm:"type:int;column:department_id;not null"`,
// 		"Title":[]string{"required", "min_cn:3", "max_cn:255"},
// 		"Type":         int    `gorm:"type:int;column:type"`,
// 		"RedirectTo":     int    `gorm:"type:int;column:redirect_to"`,
// 		"MediaId":        int    `gorm:"type:int;column:media_id;not null"`,
// 		"MediaType":      int    `gorm:"type:int;column:media_type"`,
// 		"Size":[]string{"required", "min_cn:3", "max_cn:255"},
// 		"RedirectParams":[]string{"required", "min_cn:3", "max_cn:255"},
// 		"Description":[]string{"required", "min_cn:3", "max_cn:255"},

// 	}
// 	messages := govalidator.MapData{
// 		// "name": []string{
// 		//     "required:名称为必填项",
// 		//     "min_cn:名称长度需至少 2 个字",
// 		//     "max_cn:名称长度不能超过 8 个字",
// 		//     "not_exists:名称已存在",
// 		// },
// 		// "description": []string{
// 		//     "min_cn:描述长度需至少 3 个字",
// 		//     "max_cn:描述长度不能超过 255 个字",
// 		// },
// 	}
// 	return validate(data, rules, messages)
// }
