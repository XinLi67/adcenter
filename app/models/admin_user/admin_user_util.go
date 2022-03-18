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

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (adminUserModel AdminUser) {
	database.DB.
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("user_name = ?", loginID).
		First(&adminUserModel)
	return
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (adminUserModel AdminUser) {
	database.DB.Where("phone = ?", phone).First(&adminUserModel)
	return
}
