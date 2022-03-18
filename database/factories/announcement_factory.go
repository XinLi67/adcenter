package factories

import (
	"adcenter/app/models/announcement"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeAnnouncements(count int) []announcement.Announcement {

	var objs []announcement.Announcement

	// 设置唯一性，如 Announcement 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		announcementModel := announcement.Announcement{
			//FIXME()
			AnnouncementNo: uint64(rand.Int63n(200)),
			CreatorId:      uint64(rand.Int63n(10)),
			DepartmentId:   uint64(rand.Int63n(10)),
			Title:          faker.Name(),
			LongTitle:      faker.Sentence(),
			Type:           uint64(rand.Int63n(2)),
			Banner:         faker.URL(),
			RedirectTo:     uint64(rand.Int63n(2)),
			RedirectParams: faker.Name(),
			Content:        faker.Paragraph(),
		}
		objs = append(objs, announcementModel)
	}

	return objs
}
