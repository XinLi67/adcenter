package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Post struct {
		models.BaseModel

		// Name     string `gorm:"type:varchar(255);not null;index"`
		Title string `gorm:"conlunm:title;type:varchar(255);not null;"`
		//Email    string `gorm:"type:varchar(255);index;default:null"`
		Content string `gorm:"conlumn:content;type:text;default:null;"`
		//Phone    string `gorm:"type:varchar(20);index;default:null"`
		//Password string `gorm:"type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Post{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Post{})
	}

	migrate.Add("2022_03_10_105707_create_posts_table", up, down)
}
