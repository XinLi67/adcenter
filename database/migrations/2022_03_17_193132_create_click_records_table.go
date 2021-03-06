package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func init() {

	type ClickRecord struct {
		models.BaseModel

		AdvertisingId uint64    `gorm:"type:bigint unsigned;column:advertising_id"`
		CustomerId    uint64    `gorm:"type:bigint unsigned;column:customer_id"`
		BrowsingTime  uint64    `gorm:"type:bigint unsigned;column:browsing_time"`
		StartTime     time.Time `gorm:"type:datetime(3);column:start_time"`
		EndTime       time.Time `gorm:"type:datetime(3);column:end_time"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ClickRecord{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ClickRecord{})
	}

	migrate.Add("2022_03_17_193132_create_click_records_table", up, down)
}
