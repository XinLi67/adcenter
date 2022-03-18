package migrations

import (
	"adcenter/app/models"
	"adcenter/app/models/role"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type AdminUser struct {
		models.BaseModel

		Name         string      `gorm:"type:varchar(60);column:name"`
		Username     string      `gorm:"type:varchar(60);column:user_name"`
		DepartmentId uint64      `gorm:"type:bigint unsigned;column:department_id"`
		Gender       uint64      `gorm:"type:tinyint unsigned;column:gender;"`
		Email        string      `gorm:"type:varchar(128);column:email;"`
		Phone        string      `gorm:"type:varchar(20);column:phone;"`
		Password     string      `gorm:"type:varchar(128);column:password;"`
		Status       uint64      `gorm:"type:tinyint unsigned;column:status"`
		Role         []role.Role `gorm:"many2many:user_role_rel"`

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
