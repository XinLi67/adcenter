//Package menu 模型
package menu

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Menu struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	MenuName       string `gorm:"type:varchar(60);column:menu_name"`
	ParentId       uint64 `gorm:"type:smallint unsigned;column:parent_id"`
	Icon           string `gorm:"type:varchar(60);column:icon"`
	Uri            string `gorm:"type:varchar(255);column:uri"`
	IsLink         uint64 `gorm:"type:tinyint unsigned;column:is_link"`
	PermissionName string `gorm:"type:varchar(60);column:permission_name"`
	GuardName      string `gorm:"type:varchar(60);column:guard_name"`
	Sequence       uint64 `gorm:"type:smallint unsigned;column:sequence"`

	models.CommonTimestampsField
}

func (menu *Menu) Create() {
	database.DB.Create(&menu)
}

func (menu *Menu) Save() (rowsAffected int64) {
	result := database.DB.Save(&menu)
	return result.RowsAffected
}

func (menu *Menu) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&menu)
	return result.RowsAffected
}
