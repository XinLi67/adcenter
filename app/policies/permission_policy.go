package policies

import (
	"adcenter/app/models/permission"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyPermission(c *gin.Context, permissionModel permission.Permission) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(permissionModel.ID, 10)
}

// func CanViewPermission(c *gin.Context, permissionModel permission.Permission) bool {}
// func CanCreatePermission(c *gin.Context, permissionModel permission.Permission) bool {}
// func CanUpdatePermission(c *gin.Context, permissionModel permission.Permission) bool {}
// func CanDeletePermission(c *gin.Context, permissionModel permission.Permission) bool {}
