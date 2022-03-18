package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Media struct {
		models.BaseModel

		// MediaId      uint64 `gorm:"type:bigint unsigned;column:media_id;not null"`
		CreatorId    uint64 `gorm:"type:bigint unsigned;column:creator_id"`
		MediaGroupId uint64 `gorm:"type:bigint unsigned;column:media_group_id"`
		DepartmentId uint64 `gorm:"type:bigint unsigned;column:department_id"`
		Type         uint64 `gorm:"type:tinyint unsigned;column:type"`
		Url          string `gorm:"type:varchar(255);column:url"`
		Title        string `gorm:"type:varchar(255);column:title"`
		Content      string `gorm:"type:text;column:content"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Media{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Media{})
	}

	migrate.Add("2022_03_17_144955_create_medias_table", up, down)
}
