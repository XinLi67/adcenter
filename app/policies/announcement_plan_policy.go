package policies

import (
	"adcenter/app/models/announcement_plan"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(announcementPlanModel.ID, 10)
}

// func CanViewAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
// func CanCreateAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
// func CanUpdateAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
// func CanDeleteAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
