package policies

import (
	"adcenter/app/models/media_group"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyMediaGroup(c *gin.Context, mediaGroupModel media_group.MediaGroup) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(mediaGroupModel.ID, 10)
}

// func CanViewMediaGroup(c *gin.Context, mediaGroupModel media_group.MediaGroup) bool {}
// func CanCreateMediaGroup(c *gin.Context, mediaGroupModel media_group.MediaGroup) bool {}
// func CanUpdateMediaGroup(c *gin.Context, mediaGroupModel media_group.MediaGroup) bool {}
// func CanDeleteMediaGroup(c *gin.Context, mediaGroupModel media_group.MediaGroup) bool {}
