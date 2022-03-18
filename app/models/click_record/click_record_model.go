//Package click_record 模型
package click_record

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
	"time"
)

type ClickRecord struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	AdvertisingId uint64    `gorm:"type:bigint unsigned;column:advertising_id"`
	CustomerId    uint64    `gorm:"type:bigint unsigned;column:customer_id"`
	BrowsingTime  uint64    `gorm:"type:bigint unsigned;column:browsing_time"`
	StartTime     time.Time `gorm:"type:datetime(3);column:start_time"`
	EndTime       time.Time `gorm:"type:datetime(3);column:end_time"`

	models.CommonTimestampsField
}

func (clickRecord *ClickRecord) Create() {
	database.DB.Create(&clickRecord)
}

func (clickRecord *ClickRecord) Save() (rowsAffected int64) {
	result := database.DB.Save(&clickRecord)
	return result.RowsAffected
}

func (clickRecord *ClickRecord) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&clickRecord)
	return result.RowsAffected
}
