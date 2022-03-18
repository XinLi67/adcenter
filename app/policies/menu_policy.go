package policies

import (
	"adcenter/app/models/menu"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyMenu(c *gin.Context, menuModel menu.Menu) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(menuModel.ID, 10)
}

// func CanViewMenu(c *gin.Context, menuModel menu.Menu) bool {}
// func CanCreateMenu(c *gin.Context, menuModel menu.Menu) bool {}
// func CanUpdateMenu(c *gin.Context, menuModel menu.Menu) bool {}
// func CanDeleteMenu(c *gin.Context, menuModel menu.Menu) bool {}
