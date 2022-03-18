package channel

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (channel Channel) {
    database.DB.Where("id", idstr).First(&channel)
    return
}

func GetBy(field, value string) (channel Channel) {
    database.DB.Where("? = ?", field, value).First(&channel)
    return
}

func All() (channels []Channel) {
    database.DB.Find(&channels)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Channel{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (channels []Channel, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Channel{}),
        &channels,
        app.V1URL(database.TableName(&Channel{})),
        perPage,
    )
    return
}