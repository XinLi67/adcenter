package policies

import (
	"adcenter/app/models/click_record"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(clickRecordModel.ID, 10)
}

// func CanViewClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
// func CanCreateClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
// func CanUpdateClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
// func CanDeleteClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
