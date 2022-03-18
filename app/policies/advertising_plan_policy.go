package policies

import (
	"adcenter/app/models/advertising_plan"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(advertisingPlanModel.ID, 10)
}

// func CanViewAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
// func CanCreateAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
// func CanUpdateAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
// func CanDeleteAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
