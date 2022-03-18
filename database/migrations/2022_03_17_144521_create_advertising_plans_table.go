package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func init() {

	type AdvertisingPlan struct {
		models.BaseModel

		// AdvertisingPlanId     uint64    `gorm:"type:bigint unsigned;column:advertising_plan_id;not null"`
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

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AdvertisingPlan{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AdvertisingPlan{})
	}

	migrate.Add("2022_03_17_144521_create_advertising_plans_table", up, down)
}
