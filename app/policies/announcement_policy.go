package policies

import (
	"adcenter/app/models/announcement"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(announcementModel.ID, 10)
}

// func CanViewAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
// func CanCreateAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
// func CanUpdateAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
// func CanDeleteAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
