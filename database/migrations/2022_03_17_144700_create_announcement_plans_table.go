package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func init() {

	type AnnouncementPlan struct {
		models.BaseModel

		// AnnouncementPlanId     uint64    `gorm:"type:bigint unsigned;column:announcement_plan_id;not null"`
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

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AnnouncementPlan{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AnnouncementPlan{})
	}

	migrate.Add("2022_03_17_144700_create_announcement_plans_table", up, down)
}
