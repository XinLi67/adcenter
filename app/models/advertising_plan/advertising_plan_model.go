//Package advertising_plan 模型
package advertising_plan

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
	"time"
)

type AdvertisingPlan struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	CreatorId             uint64    `gorm:"type:bigint unsigned;column:creator_id"`
	AdvertisingId         uint64    `gorm:"type:bigint unsigned;column:advertising_id"`
	AdvertisingType       uint64    `gorm:"type:tinyint unsigned;column:advertising_type"`
	AdvertisingPositionId uint64    `gorm:"type:bigint unsigned;column:advertising_position_id"`
	Order                 uint64    `gorm:"type:smallint unsigned;column:order"`
	SchedulingDate        uint64    `gorm:"type:tinyint unsigned;column:scheduling_date"`
	SchedulingTime        uint64    `gorm:"type:tinyint unsigned;column:scheduling_time"`
	StartDate             time.Time `gorm:"type:datetime(3);column:start_date"`
	EndTDate              time.Time `gorm:"type:datetime(3);column:end_date"`
	StartTime             time.Time `gorm:"type:datetime(3);column:start_time"`
	EndTime               time.Time `gorm:"type:datetime(3);column:end_time"`
	Status                uint64    `gorm:"type:tinyint unsigned;column:status"`

	models.CommonTimestampsField
}

func (advertisingPlan *AdvertisingPlan) Create() {
	database.DB.Create(&advertisingPlan)
}

func (advertisingPlan *AdvertisingPlan) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertisingPlan)
	return result.RowsAffected
}

func (advertisingPlan *AdvertisingPlan) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertisingPlan)
	return result.RowsAffected
}
