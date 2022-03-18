package factories

import (
	"adcenter/app/models/announcement_plan"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func MakeAnnouncementPlans(count int) []announcement_plan.AnnouncementPlan {

	var objs []announcement_plan.AnnouncementPlan

	// 设置唯一性，如 AnnouncementPlan 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		announcementPlanModel := announcement_plan.AnnouncementPlan{
			// FIXME()
			CreatorId:              uint64(rand.Int63n(10)),
			AnnouncementId:         uint64(rand.Int63n(10)),
			AnnouncementType:       uint64(rand.Int63n(2)),
			AnnouncementPositionId: uint64(rand.Int63n(10)),
			Order:                  uint64(rand.Int63n(20)),
			SchedulingDate:         uint64(rand.Int63n(2)),
			SchedulingTime:         uint64(rand.Int63n(2)),
			StartDate:              time.Unix(faker.UnixTime(), 0).UTC().Local(),
			EndTDate:               time.Unix(faker.UnixTime(), 0).UTC().Local(),
			StartTime:              time.Unix(faker.UnixTime(), 0).UTC().Local(),
			EndTime:                time.Unix(faker.UnixTime(), 0).UTC().Local(),
			Status:                 uint64(rand.Int63n(2)),
		}
		objs = append(objs, announcementPlanModel)
	}

	return objs
}
