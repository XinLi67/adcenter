//Package announcement_plan 模型
package announcement_plan

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
	"time"
)

type AnnouncementPlan struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	CreatorId              uint64    `gorm:"type:bigint unsigned;column:creator_id"`
	AnnouncementId         uint64    `gorm:"type:bigint unsigned;column:announcement_id"`
	AnnouncementType       uint64    `gorm:"type:tinyint unsigned;column:announcement_type"`
	AnnouncementPositionId uint64    `gorm:"type:bigint unsigned;column:announcement_position_id"`
	Order                  uint64    `gorm:"type:smallint unsigned;column:order"`
	SchedulingDate         uint64    `gorm:"type:tinyint unsigned;column:scheduling_date"`
	SchedulingTime         uint64    `gorm:"type:tinyint unsigned;column:scheduling_time"`
	StartDate              time.Time `gorm:"type:datetime(3);column:start_date"`
	EndTDate               time.Time `gorm:"type:datetime(3);column:end_date"`
	StartTime              time.Time `gorm:"type:datetime(3);column:start_time"`
	EndTime                time.Time `gorm:"type:datetime(3);column:end_time"`
	Status                 uint64    `gorm:"type:tinyint unsigned;column:status"`

	models.CommonTimestampsField
}

func (announcementPlan *AnnouncementPlan) Create() {
	database.DB.Create(&announcementPlan)
}

func (announcementPlan *AnnouncementPlan) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcementPlan)
	return result.RowsAffected
}

func (announcementPlan *AnnouncementPlan) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&announcementPlan)
	return result.RowsAffected
}
