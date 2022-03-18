package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type AdvertisingPosition struct {
		models.BaseModel

		// AdvertisingPositionId   uint64 `gorm:"type:bigint unsigned;column:advertising_position_id;not null"`
		AdvertisingPositionName string `gorm:"type:varchar(60);column:advertising_position_name"`
		ParentId                uint64 `gorm:"type:smallint unsigned;column:parent_id"`
		Code                    string `gorm:"type:varchar(60);column:code"`
		Height                  uint64 `gorm:"type:int unsigned;column:height"`
		Weight                  uint64 `gorm:"type:int unsigned;column:weight"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AdvertisingPosition{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AdvertisingPosition{})
	}

	migrate.Add("2022_03_17_144015_create_advertising_positions_table", up, down)
}
