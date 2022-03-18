package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Menu struct {
		models.BaseModel

		// MenuId         uint64 `gorm:"type:bigint unsigned;column:menu_id;not null"`
		MenuName       string `gorm:"type:varchar(60);column:menu_name"`
		ParentId       uint64 `gorm:"type:smallint unsigned;column:parent_id"`
		Icon           string `gorm:"type:varchar(60);column:icon"`
		Uri            string `gorm:"type:varchar(255);column:uri"`
		IsLink         uint64 `gorm:"type:tinyint unsigned;column:is_link"`
		PermissionName string `gorm:"type:varchar(60);column:permission_name"`
		GuardName      string `gorm:"type:varchar(60);column:guard_name"`
		Sequence       uint64 `gorm:"type:smallint unsigned;column:sequence"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Menu{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Menu{})
	}

	migrate.Add("2022_03_17_150452_create_menus_table", up, down)
}
