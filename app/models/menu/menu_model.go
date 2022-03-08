//Package menu 模型
package menu

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Menu struct {
	models.BaseModel

	ParentId  int64  `gorm:"column:parent_id;type:bigint unsigned;not null;index"`
	Icon      string `gorm:"column:icon;type:varchar(255);null"`
	Uri       string `gorm:"column:uri;type:varchar(255);null"`
	Name      string `gorm:"column:name;type:varchar(60);not null;index"`
	GuardName string `gorm:"column:guard_name;type:varchar(60);not null;default:admin"`
	Sequence  int64  `gorm:"column:sequence;type:int;null"`

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
