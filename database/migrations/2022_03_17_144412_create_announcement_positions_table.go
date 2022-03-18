package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type AnnouncementPosition struct {
		models.BaseModel

		// AnnouncementPositionId   uint64 `gorm:"type:bigint unsigned;column:announcement_position_id;not null"`
		AnnouncementPositionName string `gorm:"type:varchar(60);column:announcement_position_name"`
		ParentId                 uint64 `gorm:"type:smallint unsigned;column:parent_id"`
		Code                     string `gorm:"type:varchar(60);column:code"`
		Height                   uint64 `gorm:"type:int unsigned;column:height"`
		Weight                   uint64 `gorm:"type:int unsigned;column:weight"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AnnouncementPosition{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AnnouncementPosition{})
	}

	migrate.Add("2022_03_17_144412_create_announcement_positions_table", up, down)
}
