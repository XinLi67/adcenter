package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Permission struct {
		models.BaseModel

		// PermissionId      uint64 `gorm:"type:bigint unsigned;column:permission_id;not null"`
		PermissionGroupId uint64 `gorm:"type:bigint unsigned;column:permission_group_id"`
		PermissionName    string `gorm:"type:varchar(60);column:permission_name"`
		Icon              string `gorm:"type:varchar(60);column:icon"`
		GuardName         string `gorm:"type:varchar(60);column:guard_name"`
		DisplayName       string `gorm:"type:varchar(60);column:display_name"`
		Sequence          uint64 `gorm:"type:smallint unsigned;column:sequence"`
		Description       string `gorm:"type:text;column:description"`
		CreatedName       string `gorm:"type:varchar(60);column:created_name"`
		UpdatedName       string `gorm:"type:varchar(60);column:updated_name"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Permission{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Permission{})
	}

	migrate.Add("2022_03_17_150119_create_permissions_table", up, down)
}
