//Package admin_user 模型
package admin_user

import (
	"adcenter/app/models"
	"adcenter/app/models/role"
	"adcenter/pkg/database"
)

type AdminUser struct {
	models.BaseModel

	Name         string      `gorm:"type:varchar(60);column:name"`
	Username     string      `gorm:"type:varchar(60);column:user_name"`
	DepartmentId uint64      `gorm:"type:bigint unsigned;column:department_id"`
	Gender       uint64      `gorm:"type:tinyint unsigned;column:gender;"`
	Email        string      `gorm:"type:varchar(128);column:email;"`
	Phone        string      `gorm:"type:varchar(20);column:phone;"`
	Password     string      `gorm:"type:varchar(128)column:password;"`
	Status       uint64      `gorm:"type:tinyint unsigned;column:status"`
	Role         []role.Role `gorm:"many2many:user_role_rel"`

	models.CommonTimestampsField
}

func (adminUser *AdminUser) Create() {
	database.DB.Create(&adminUser)
}

func (adminUser *AdminUser) Save() (rowsAffected int64) {
	result := database.DB.Save(&adminUser)
	return result.RowsAffected
}

func (adminUser *AdminUser) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&adminUser)
	return result.RowsAffected
}
