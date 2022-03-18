//Package advertising_position 模型
package advertising_position

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type AdvertisingPosition struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	AdvertisingPositionName string `gorm:"type:varchar(60);column:advertising_position_name"`
	ParentId                uint64 `gorm:"type:smallint unsigned;column:parent_id"`
	Code                    string `gorm:"type:varchar(60);column:code"`
	Height                  uint64 `gorm:"type:int unsigned;column:height"`
	Weight                  uint64 `gorm:"type:int unsigned;column:weight"`

	models.CommonTimestampsField
}

func (advertisingPosition *AdvertisingPosition) Create() {
	database.DB.Create(&advertisingPosition)
}

func (advertisingPosition *AdvertisingPosition) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertisingPosition)
	return result.RowsAffected
}

func (advertisingPosition *AdvertisingPosition) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertisingPosition)
	return result.RowsAffected
}
