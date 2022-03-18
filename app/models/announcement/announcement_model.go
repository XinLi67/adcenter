//Package announcement 模型
package announcement

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Announcement struct {
	models.BaseModel

	// Put fields in here
	//FIXME()
	AnnouncementNo uint64 `gorm:"type:bigint unsigned;column:announcement_no"`
	CreatorId      uint64 `gorm:"type:bigint unsigned;column:creator_id"`
	DepartmentId   uint64 `gorm:"type:bigint unsigned;column:department_id"`
	Title          string `gorm:"type:varchar(255);column:title"`
	LongTitle      string `gorm:"type:varchar(255);column:long_title"`
	Type           uint64 `gorm:"type:tinyint unsigned;column:type"`
	Banner         string `gorm:"type:varchar(255);column:banner"`
	RedirectTo     uint64 `gorm:"type:tinyint unsigned;column:redirect_to"`
	RedirectParams string `gorm:"type:varchar(60);column:redirect_params"`
	Content        string `gorm:"type:text;column:content"`

	models.CommonTimestampsField
}

func (announcement *Announcement) Create() {
	database.DB.Create(&announcement)
}

func (announcement *Announcement) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcement)
	return result.RowsAffected
}

func (announcement *Announcement) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&announcement)
	return result.RowsAffected
}
