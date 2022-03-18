//Package permission_group 模型
package permission_group

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type PermissionGroup struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	PermissionGroupName string `gorm:"type:varchar(60);column:permission_group_name"`

	models.CommonTimestampsField
}

func (permissionGroup *PermissionGroup) Create() {
	database.DB.Create(&permissionGroup)
}

func (permissionGroup *PermissionGroup) Save() (rowsAffected int64) {
	result := database.DB.Save(&permissionGroup)
	return result.RowsAffected
}

func (permissionGroup *PermissionGroup) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&permissionGroup)
	return result.RowsAffected
}
