//Package media 模型
package media

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Media struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	CreatorId    uint64 `gorm:"type:bigint unsigned;column:creator_id"`
	MediaGroupId uint64 `gorm:"type:bigint unsigned;column:media_group_id"`
	DepartmentId uint64 `gorm:"type:bigint unsigned;column:department_id"`
	Type         uint64 `gorm:"type:tinyint unsigned;column:type"`
	Url          string `gorm:"type:varchar(255);column:url"`
	Title        string `gorm:"type:varchar(255);column:title"`
	Content      string `gorm:"type:text;column:content"`

	models.CommonTimestampsField
}

func (media *Media) Create() {
	database.DB.Create(&media)
}

func (media *Media) Save() (rowsAffected int64) {
	result := database.DB.Save(&media)
	return result.RowsAffected
}

func (media *Media) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&media)
	return result.RowsAffected
}
