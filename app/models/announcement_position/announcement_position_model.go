//Package announcement_position 模型
package announcement_position

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type AnnouncementPosition struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	AnnouncementPositionName string `gorm:"type:varchar(60);column:announcement_position_name"`
	ParentId                 uint64 `gorm:"type:smallint unsigned;column:parent_id"`
	Code                     string `gorm:"type:varchar(60);column:code"`
	Height                   uint64 `gorm:"type:int unsigned;column:height"`
	Weight                   uint64 `gorm:"type:int unsigned;column:weight"`

	models.CommonTimestampsField
}

func (announcementPosition *AnnouncementPosition) Create() {
	database.DB.Create(&announcementPosition)
}

func (announcementPosition *AnnouncementPosition) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcementPosition)
	return result.RowsAffected
}

func (announcementPosition *AnnouncementPosition) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&announcementPosition)
	return result.RowsAffected
}
