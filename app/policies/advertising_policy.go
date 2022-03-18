package policies

import (
	"adcenter/app/models/advertising"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(advertisingModel.ID, 10)
}

// func CanViewAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
// func CanCreateAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
// func CanUpdateAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
// func CanDeleteAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
