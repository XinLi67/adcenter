//Package media_group 模型
package media_group

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type MediaGroup struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	MediaGroupName string `gorm:"type:varchar(60);column:media_group_name"`

	models.CommonTimestampsField
}

func (mediaGroup *MediaGroup) Create() {
	database.DB.Create(&mediaGroup)
}

func (mediaGroup *MediaGroup) Save() (rowsAffected int64) {
	result := database.DB.Save(&mediaGroup)
	return result.RowsAffected
}

func (mediaGroup *MediaGroup) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&mediaGroup)
	return result.RowsAffected
}
