package announcement_plan

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (announcementPlan AnnouncementPlan) {
    database.DB.Where("id", idstr).First(&announcementPlan)
    return
}

func GetBy(field, value string) (announcementPlan AnnouncementPlan) {
    database.DB.Where("? = ?", field, value).First(&announcementPlan)
    return
}

func All() (announcementPlans []AnnouncementPlan) {
    database.DB.Find(&announcementPlans)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(AnnouncementPlan{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (announcementPlans []AnnouncementPlan, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(AnnouncementPlan{}),
        &announcementPlans,
        app.V1URL(database.TableName(&AnnouncementPlan{})),
        perPage,
    )
    return
}