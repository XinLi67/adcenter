//Package admin_user 模型
package admin_user

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type AdminUser struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(60);not null;index"`
	Username string `gorm:"type:varchar(60);not null"`
	Email    string `gorm:"type:varchar(255);index;default:null"`
	Phone    string `gorm:"type:varchar(20);index;default:null"`
	Password string `gorm:"type:varchar(255)"`

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
