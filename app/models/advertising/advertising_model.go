//Package advertising 模型
package advertising

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Advertising struct {
	models.BaseModel

	// Put fields in here
	//FIXME()
	AdvertisingNo  uint64 `gorm:"type:bigint unsigned;column:advertising_no"`
	CreatorId      uint64 `gorm:"type:bigint unsigned;column:creator_id"`
	DepartmentId   uint64 `gorm:"type:bigint unsigned;column:department_id"`
	Title          string `gorm:"type:varchar(255);column:title"`
	Type           uint64 `gorm:"type:tinyint unsigned;column:type"`
	RedirectTo     uint64 `gorm:"type:tinyint unsigned;column:redirect_to"`
	MediaId        uint64 `gorm:"type:bigint unsigned;column:media_id"`
	MediaType      uint64 `gorm:"type:tinyint unsigned;column:media_type"`
	Size           string `gorm:"type:varchar(60);column:size"`
	RedirectParams string `gorm:"type:varchar(60);column:redirect_params"`
	Description    string `gorm:"type:text;column:description"`

	models.CommonTimestampsField
}

func (advertising *Advertising) Create() {
	database.DB.Create(&advertising)
}

func (advertising *Advertising) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertising)
	return result.RowsAffected
}

func (advertising *Advertising) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertising)
	return result.RowsAffected
}
