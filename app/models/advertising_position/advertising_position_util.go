package advertising_position

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (advertisingPosition AdvertisingPosition) {
    database.DB.Where("id", idstr).First(&advertisingPosition)
    return
}

func GetBy(field, value string) (advertisingPosition AdvertisingPosition) {
    database.DB.Where("? = ?", field, value).First(&advertisingPosition)
    return
}

func All() (advertisingPositions []AdvertisingPosition) {
    database.DB.Find(&advertisingPositions)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(AdvertisingPosition{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (advertisingPositions []AdvertisingPosition, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(AdvertisingPosition{}),
        &advertisingPositions,
        app.V1URL(database.TableName(&AdvertisingPosition{})),
        perPage,
    )
    return
}