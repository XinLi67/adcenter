package policies

import (
	"adcenter/app/models/media"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyMedia(c *gin.Context, mediaModel media.Media) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(mediaModel.ID, 10)
}

// func CanViewMedia(c *gin.Context, mediaModel media.Media) bool {}
// func CanCreateMedia(c *gin.Context, mediaModel media.Media) bool {}
// func CanUpdateMedia(c *gin.Context, mediaModel media.Media) bool {}
// func CanDeleteMedia(c *gin.Context, mediaModel media.Media) bool {}
