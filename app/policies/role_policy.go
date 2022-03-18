package policies

import (
	"adcenter/app/models/role"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyRole(c *gin.Context, roleModel role.Role) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(roleModel.ID, 10)
}

// func CanViewRole(c *gin.Context, roleModel role.Role) bool {}
// func CanCreateRole(c *gin.Context, roleModel role.Role) bool {}
// func CanUpdateRole(c *gin.Context, roleModel role.Role) bool {}
// func CanDeleteRole(c *gin.Context, roleModel role.Role) bool {}
