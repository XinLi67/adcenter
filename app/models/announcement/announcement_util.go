package announcement

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (announcement Announcement) {
    database.DB.Where("id", idstr).First(&announcement)
    return
}

func GetBy(field, value string) (announcement Announcement) {
    database.DB.Where("? = ?", field, value).First(&announcement)
    return
}

func All() (announcements []Announcement) {
    database.DB.Find(&announcements)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Announcement{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (announcements []Announcement, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Announcement{}),
        &announcements,
        app.V1URL(database.TableName(&Announcement{})),
        perPage,
    )
    return
}