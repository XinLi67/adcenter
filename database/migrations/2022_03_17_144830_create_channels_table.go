package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Channel struct {
		models.BaseModel

		// ChannelId   uint64 `gorm:"type:bigint unsigned;column:channel_id;not null"`
		ChannelName string `gorm:"type:varchar(60);column:channel_name"`
		Status      uint64 `gorm:"type:tinyint unsigned;column:status"`
		Description string `gorm:"type:text;column:description"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Channel{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Channel{})
	}

	migrate.Add("2022_03_17_144830_create_channels_table", up, down)
}
