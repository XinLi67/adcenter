package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type MediaGroup struct {
		models.BaseModel

		// MediaGroupId   uint64 `gorm:"type:bigint unsigned;column:media_group_id;not null"`
		MediaGroupName string `gorm:"type:varchar(60);column:media_group_name"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&MediaGroup{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&MediaGroup{})
	}

	migrate.Add("2022_03_17_145219_create_media_groups_table", up, down)
}
