//Package channel 模型
package channel

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Channel struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	ChannelName string `gorm:"type:varchar(60);column:channel_name"`
	Status      uint64 `gorm:"type:tinyint unsigned;column:status"`
	Description string `gorm:"type:text;column:description"`

	models.CommonTimestampsField
}

func (channel *Channel) Create() {
	database.DB.Create(&channel)
}

func (channel *Channel) Save() (rowsAffected int64) {
	result := database.DB.Save(&channel)
	return result.RowsAffected
}

func (channel *Channel) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&channel)
	return result.RowsAffected
}
