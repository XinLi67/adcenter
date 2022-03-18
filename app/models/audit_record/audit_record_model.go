//Package audit_record 模型
package audit_record

import (
	"adcenter/app/models"
	"adcenter/pkg/database"
)

type AuditRecord struct {
	models.BaseModel

	// Put fields in here
	// FIXME()
	AuditableId   uint64 `gorm:"type:tinyint unsigned;column:auditable_id"`
	AuditableType uint64 `gorm:"type:tinyint unsigned;column:auditable_type"`
	ApplicantId   uint64 `gorm:"type:bigint unsigned;column:applicant_id"`
	AuditorId     uint64 `gorm:"type:bigint unsigned;column:auditor_id"`
	Status        uint64 `gorm:"type:tinyint unsigned;column:status"`
	Content       string `gorm:"type:text;column:content"`

	models.CommonTimestampsField
}

func (auditRecord *AuditRecord) Create() {
	database.DB.Create(&auditRecord)
}

func (auditRecord *AuditRecord) Save() (rowsAffected int64) {
	result := database.DB.Save(&auditRecord)
	return result.RowsAffected
}

func (auditRecord *AuditRecord) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&auditRecord)
	return result.RowsAffected
}
