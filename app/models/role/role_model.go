//Package role 模型
package role

import (
	"adcenter/app/models"
	"adcenter/app/models/permission"
	"adcenter/pkg/database"
)

type Role struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	RoleName    string                  `gorm:"type:varchar(60);column:role_name"`
	GuardName   string                  `gorm:"type:varchar(60);column:guard_name"`
	Description string                  `gorm:"type:text;column:description"`
	Permission  []permission.Permission `gorm:"many2many:role_permission_rel"`

	models.CommonTimestampsField
}

func (role *Role) Create() {
	database.DB.Create(&role)
}

func (role *Role) Save() (rowsAffected int64) {
	result := database.DB.Save(&role)
	return result.RowsAffected
}

func (role *Role) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&role)
	return result.RowsAffected
}
