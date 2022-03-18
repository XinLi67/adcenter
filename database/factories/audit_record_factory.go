package factories

import (
	"adcenter/app/models/audit_record"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeAuditRecords(count int) []audit_record.AuditRecord {

	var objs []audit_record.AuditRecord

	// 设置唯一性，如 AuditRecord 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		auditRecordModel := audit_record.AuditRecord{
			// FIXME()
			AuditableId:   uint64(rand.Int63n(2)),
			AuditableType: uint64(rand.Int63n(4)),
			ApplicantId:   uint64(rand.Int63n(10)),
			AuditorId:     uint64(rand.Int63n(10)),
			Status:        uint64(rand.Int63n(2)),
			Content:       faker.Sentence(),
		}
		objs = append(objs, auditRecordModel)
	}

	return objs
}
