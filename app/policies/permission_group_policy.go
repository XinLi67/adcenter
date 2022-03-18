package policies

import (
	"adcenter/app/models/permission_group"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyPermissionGroup(c *gin.Context, permissionGroupModel permission_group.PermissionGroup) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(permissionGroupModel.ID, 10)
}

// func CanViewPermissionGroup(c *gin.Context, permissionGroupModel permission_group.PermissionGroup) bool {}
// func CanCreatePermissionGroup(c *gin.Context, permissionGroupModel permission_group.PermissionGroup) bool {}
// func CanUpdatePermissionGroup(c *gin.Context, permissionGroupModel permission_group.PermissionGroup) bool {}
// func CanDeletePermissionGroup(c *gin.Context, permissionGroupModel permission_group.PermissionGroup) bool {}
