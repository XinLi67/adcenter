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

		ParentId  int64  `gorm:"column:parent_id;type:bigint unsigned;not null;index"`
		Icon      string `gorm:"column:icon;type:varchar(255);null"`
		Uri       string `gorm:"column:uri;type:varchar(255);null"`
		Name      string `gorm:"column:name;type:varchar(60);not null;index"`
		GuardName string `gorm:"column:guard_name;type:varchar(60);not null;default:admin"`
		Sequence  int64  `gorm:"column:sequence;type:int;null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Menu{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Menu{})
	}

	migrate.Add("2022_03_08_171914_create_menus_table", up, down)
}
