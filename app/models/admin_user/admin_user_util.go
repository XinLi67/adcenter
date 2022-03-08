package admin_user

import (
	"adcenter/pkg/app"
	"adcenter/pkg/database"
	"adcenter/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (adminUser AdminUser) {
    database.DB.Where("id", idstr).First(&adminUser)
    return
}

func GetBy(field, value string) (adminUser AdminUser) {
    database.DB.Where("? = ?", field, value).First(&adminUser)
    return
}

func All() (adminUsers []AdminUser) {
    database.DB.Find(&adminUsers)
    return 
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(AdminUser{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (adminUsers []AdminUser, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(AdminUser{}),
        &adminUsers,
        app.V1URL(database.TableName(&AdminUser{})),
        perPage,
    )
    return
}