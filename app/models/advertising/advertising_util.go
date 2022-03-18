package advertising

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (advertising Advertising) {
    database.DB.Where("id", idstr).First(&advertising)
    return
}

func GetBy(field, value string) (advertising Advertising) {
    database.DB.Where("? = ?", field, value).First(&advertising)
    return
}

func All() (advertisings []Advertising) {
    database.DB.Find(&advertisings)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Advertising{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (advertisings []Advertising, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Advertising{}),
        &advertisings,
        app.V1URL(database.TableName(&Advertising{})),
        perPage,
    )
    return
}