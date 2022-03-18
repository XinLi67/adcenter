package policies

import (
	"adcenter/app/models/advertising_position"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(advertisingPositionModel.ID, 10)
}

// func CanViewAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
// func CanCreateAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
// func CanUpdateAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
// func CanDeleteAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
