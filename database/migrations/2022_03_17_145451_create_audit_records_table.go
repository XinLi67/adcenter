package migrations

import (
	"adcenter/app/models"
	"adcenter/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type AuditRecord struct {
		models.BaseModel

		// AuditRecordId uint64 `gorm:"type:bigint unsigned;column:audit_record_id;not null"`
		AuditableId   uint64 `gorm:"type:tinyint unsigned;column:auditable_id"`
		AuditableType uint64 `gorm:"type:tinyint unsigned;column:auditable_type"`
		ApplicantId   uint64 `gorm:"type:bigint unsigned;column:applicant_id"`
		AuditorId     uint64 `gorm:"type:bigint unsigned;column:auditor_id"`
		Status        uint64 `gorm:"type:tinyint unsigned;column:status"`
		Content       string `gorm:"type:text;column:content"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AuditRecord{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AuditRecord{})
	}

	migrate.Add("2022_03_17_145451_create_audit_records_table", up, down)
}
