package migrations

import (
	"adcenter/app/models"
	"adcenter/app/models/permission"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Role struct {
		models.BaseModel

		// RoleId      uint64 `gorm:"type:bigint unsigned;column:role_id;not null"`
		RoleName    string                  `gorm:"type:varchar(60);column:role_name"`
		GuardName   string                  `gorm:"type:varchar(60);column:guard_name"`
		Description string                  `gorm:"type:text;column:description"`
		Permission  []permission.Permission `gorm:"many2many:role_permission_rel"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Role{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Role{})
	}

	migrate.Add("2022_03_17_145709_create_roles_table", up, down)
}
