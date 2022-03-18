package media

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (media Media) {
    database.DB.Where("id", idstr).First(&media)
    return
}

func GetBy(field, value string) (media Media) {
    database.DB.Where("? = ?", field, value).First(&media)
    return
}

func All() (media []Media) {
    database.DB.Find(&media)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Media{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (media []Media, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Media{}),
        &media,
        app.V1URL(database.TableName(&Media{})),
        perPage,
    )
    return
}