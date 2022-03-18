package factories

import (
	"adcenter/app/models/media"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeMedia(count int) []media.Media {

	var objs []media.Media

	// 设置唯一性，如 Media 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		mediaModel := media.Media{
			// FIXME()
			CreatorId:    uint64(rand.Int63n(10)),
			MediaGroupId: uint64(rand.Int63n(10)),
			DepartmentId: uint64(rand.Int63n(10)),
			Type:         uint64(rand.Int63n(2)),
			Url:          faker.URL(),
			Title:        faker.Name(),
			Content:      faker.Sentence(),
		}
		objs = append(objs, mediaModel)
	}

	return objs
}
