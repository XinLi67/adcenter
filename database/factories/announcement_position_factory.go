package factories

import (
	"adcenter/app/models/announcement_position"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeAnnouncementPositions(count int) []announcement_position.AnnouncementPosition {

	var objs []announcement_position.AnnouncementPosition

	// 设置唯一性，如 AnnouncementPosition 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		announcementPositionModel := announcement_position.AnnouncementPosition{
			// FIXME()
			AnnouncementPositionName: faker.Name(),
			ParentId:                 uint64(rand.Int63n(20)),
			Code:                     faker.Name(),
			Height:                   uint64(rand.Int63n(1000)),
			Weight:                   uint64(rand.Int63n(1000)),
		}
		objs = append(objs, announcementPositionModel)
	}

	return objs
}
