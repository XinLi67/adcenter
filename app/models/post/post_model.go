//Package post 模型
package post

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type Post struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	//这里需要添加新的内容
	Title   string `gorm:"conlunm:title;type:varchar(255);not null;"`
	Content string `gorm:"conlumn:content;type:text;default:null;"`

	models.CommonTimestampsField
}

func (post *Post) Create() {
	database.DB.Create(&post)
}

func (post *Post) Save() (rowsAffected int64) {
	result := database.DB.Save(&post)
	return result.RowsAffected
}

func (post *Post) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&post)
	return result.RowsAffected
}
