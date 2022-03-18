package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Department struct {
		models.BaseModel

		// DepartmentId   uint64 `gorm:"type:bigint unsigned;column:department_id;not null"`
		DepartmentName string `gorm:"type:varchar(60);column:department_name"`
		ParentId       uint64 `gorm:"type:smallint unsigned;column:parent_id"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Department{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Department{})
	}

	migrate.Add("2022_03_17_145610_create_departments_table", up, down)
}
