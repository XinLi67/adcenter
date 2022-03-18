package factories

import (
	"adcenter/app/models/media_group"

	"github.com/bxcodec/faker/v3"
)

func MakeMediaGroups(count int) []media_group.MediaGroup {

	var objs []media_group.MediaGroup

	// 设置唯一性，如 MediaGroup 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		mediaGroupModel := media_group.MediaGroup{
			// FIXME()
			MediaGroupName: faker.Name(),
		}
		objs = append(objs, mediaGroupModel)
	}

	return objs
}
