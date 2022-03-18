package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type PermissionGroup struct {
		models.BaseModel

		// PermissionGroupId   uint64 `gorm:"type:bigint unsigned;column:permission_group_id;not null"`
		PermissionGroupName string `gorm:"type:varchar(60);column:permission_group_name"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&PermissionGroup{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&PermissionGroup{})
	}

	migrate.Add("2022_03_17_150343_create_permission_groups_table", up, down)
}
