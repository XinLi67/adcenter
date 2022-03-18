package factories

import (
	"adcenter/app/models/click_record"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func MakeClickRecords(count int) []click_record.ClickRecord {

	var objs []click_record.ClickRecord

	// 设置唯一性，如 ClickRecord 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		clickRecordModel := click_record.ClickRecord{
			// FIXME()
			AdvertisingId: uint64(rand.Int63n(10)),
			CustomerId:    uint64(rand.Int63n(10)),
			BrowsingTime:  uint64(rand.Int63n(1000)),
			StartTime:     time.Unix(faker.UnixTime(), 0).UTC().Local(),
			EndTime:       time.Unix(faker.UnixTime(), 0).UTC().Local(),
		}
		objs = append(objs, clickRecordModel)
	}

	return objs
}
