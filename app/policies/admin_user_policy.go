package policies

import (
	"adcenter/app/models/admin_user"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAdminUser(c *gin.Context, adminUserModel admin_user.AdminUser) bool {
	// 这里需要转换数据类型，将unit64类型转换为string类型
	return auth.CurrentUID(c) == strconv.FormatUint(adminUserModel.ID, 10)
}

// func CanViewAdminUser(c *gin.Context, adminUserModel admin_user.AdminUser) bool {}
// func CanCreateAdminUser(c *gin.Context, adminUserModel admin_user.AdminUser) bool {}
// func CanUpdateAdminUser(c *gin.Context, adminUserModel admin_user.AdminUser) bool {}
// func CanDeleteAdminUser(c *gin.Context, adminUserModel admin_user.AdminUser) bool {}
