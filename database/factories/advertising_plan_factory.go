package factories

import (
	"adcenter/app/models/advertising_plan"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func MakeAdvertisingPlans(count int) []advertising_plan.AdvertisingPlan {

	var objs []advertising_plan.AdvertisingPlan

	// 设置唯一性，如 AdvertisingPlan 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		advertisingPlanModel := advertising_plan.AdvertisingPlan{
			// FIXME()
			CreatorId:             uint64(rand.Int63n(10)),
			AdvertisingId:         uint64(rand.Int63n(10)),
			AdvertisingType:       uint64(rand.Int63n(4)),
			AdvertisingPositionId: uint64(rand.Int63n(10)),
			Order:                 uint64(rand.Int63n(20)),
			SchedulingDate:        uint64(rand.Int63n(2)),
			SchedulingTime:        uint64(rand.Int63n(2)),
			StartDate:             time.Unix(faker.UnixTime(), 0).UTC().Local(),
			EndTDate:              time.Unix(faker.UnixTime(), 0).UTC().Local(),
			StartTime:             time.Unix(faker.UnixTime(), 0).UTC().Local(),
			EndTime:               time.Unix(faker.UnixTime(), 0).UTC().Local(),
			Status:                uint64(rand.Int63n(2)),
		}
		objs = append(objs, advertisingPlanModel)
	}

	return objs
}
