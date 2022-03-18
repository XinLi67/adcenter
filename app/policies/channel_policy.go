package policies

import (
	"adcenter/app/models/channel"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyChannel(c *gin.Context, channelModel channel.Channel) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(channelModel.ID, 10)
}

// func CanViewChannel(c *gin.Context, channelModel channel.Channel) bool {}
// func CanCreateChannel(c *gin.Context, channelModel channel.Channel) bool {}
// func CanUpdateChannel(c *gin.Context, channelModel channel.Channel) bool {}
// func CanDeleteChannel(c *gin.Context, channelModel channel.Channel) bool {}
