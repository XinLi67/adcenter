package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Advertising struct {
		models.BaseModel

		AdvertisingNo  uint64 `gorm:"type:bigint unsigned;column:advertising_no"`
		CreatorId      uint64 `gorm:"type:bigint unsigned;column:creator_id"`
		DepartmentId   uint64 `gorm:"type:bigint unsigned;column:department_id"`
		Title          string `gorm:"type:varchar(255);column:title"`
		Type           uint64 `gorm:"type:tinyint unsigned;column:type"`
		RedirectTo     uint64 `gorm:"type:tinyint unsigned;column:redirect_to"`
		MediaId        uint64 `gorm:"type:bigint unsigned;column:media_id"`
		MediaType      uint64 `gorm:"type:tinyint unsigned;column:media_type"`
		Size           string `gorm:"type:varchar(60);column:size"`
		RedirectParams string `gorm:"type:varchar(60);column:redirect_params"`
		Description    string `gorm:"type:text;column:description"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Advertising{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Advertising{})
	}

	migrate.Add("2022_03_17_143440_create_advertisings_table", up, down)
}
