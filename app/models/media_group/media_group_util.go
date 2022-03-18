package media_group

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (mediaGroup MediaGroup) {
    database.DB.Where("id", idstr).First(&mediaGroup)
    return
}

func GetBy(field, value string) (mediaGroup MediaGroup) {
    database.DB.Where("? = ?", field, value).First(&mediaGroup)
    return
}

func All() (mediaGroups []MediaGroup) {
    database.DB.Find(&mediaGroups)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(MediaGroup{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (mediaGroups []MediaGroup, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(MediaGroup{}),
        &mediaGroups,
        app.V1URL(database.TableName(&MediaGroup{})),
        perPage,
    )
    return
}