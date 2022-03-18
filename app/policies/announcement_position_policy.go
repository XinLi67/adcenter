package policies

import (
	"adcenter/app/models/announcement_position"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(announcementPositionModel.ID, 10)
}

// func CanViewAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
// func CanCreateAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
// func CanUpdateAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
// func CanDeleteAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
