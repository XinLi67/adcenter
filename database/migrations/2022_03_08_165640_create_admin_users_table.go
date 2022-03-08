package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type AdminUser struct {
		models.BaseModel

		Name     string `gorm:"column:name;type:varchar(60);not null;index"`
		Username string `gorm:"column:username;type:varchar(60);not null"`
		Email    string `gorm:"column:email;type:varchar(255);index;default:null"`
		Phone    string `gorm:"column:phone;type:varchar(20);index;default:null"`
		Password string `gorm:"column:password;type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AdminUser{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AdminUser{})
	}

	migrate.Add("2022_03_08_165640_create_admin_users_table", up, down)
}
