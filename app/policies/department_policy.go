package policies

import (
	"adcenter/app/models/department"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyDepartment(c *gin.Context, departmentModel department.Department) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(departmentModel.ID, 10)
}

// func CanViewDepartment(c *gin.Context, departmentModel department.Department) bool {}
// func CanCreateDepartment(c *gin.Context, departmentModel department.Department) bool {}
// func CanUpdateDepartment(c *gin.Context, departmentModel department.Department) bool {}
// func CanDeleteDepartment(c *gin.Context, departmentModel department.Department) bool {}
