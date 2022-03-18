package policies

import (
	"adcenter/app/models/audit_record"
	"adcenter/pkg/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CanModifyAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {
	return auth.CurrentUID(c) == strconv.FormatUint(auditRecordModel.ID, 10)
}

// func CanViewAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
// func CanCreateAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
// func CanUpdateAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
// func CanDeleteAuditRecord(c *gin.Context, auditRecordModel audit_record.AuditRecord) bool {}
