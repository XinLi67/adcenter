//Package permission 模型
package permission

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Permission struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	PermissionGroupId uint64 `gorm:"type:bigint unsigned;column:permission_group_id"`
	PermissionName    string `gorm:"type:varchar(60);column:permission_name"`
	Icon              string `gorm:"type:varchar(60);column:icon"`
	GuardName         string `gorm:"type:varchar(60);column:guard_name"`
	DisplayName       string `gorm:"type:varchar(60);column:display_name"`
	Sequence          uint64 `gorm:"type:smallint unsigned;column:sequence"`
	Description       string `gorm:"type:text;column:description"`
	CreatedName       string `gorm:"type:varchar(60);column:created_name"`
	UpdatedName       string `gorm:"type:varchar(60);column:updated_name"`

	models.CommonTimestampsField
}

func (permission *Permission) Create() {
	database.DB.Create(&permission)
}

func (permission *Permission) Save() (rowsAffected int64) {
	result := database.DB.Save(&permission)
	return result.RowsAffected
}

func (permission *Permission) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&permission)
	return result.RowsAffected
}
