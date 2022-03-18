//Package department 模型
package department

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Department struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	DepartmentName string `gorm:"type:varchar(60);column:department_name"`
	ParentId       uint64 `gorm:"type:smallint unsigned;column:parent_id"`

	models.CommonTimestampsField
}

func (department *Department) Create() {
	database.DB.Create(&department)
}

func (department *Department) Save() (rowsAffected int64) {
	result := database.DB.Save(&department)
	return result.RowsAffected
}

func (department *Department) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&department)
	return result.RowsAffected
}
