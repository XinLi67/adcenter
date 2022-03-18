package factories

import (
	"adcenter/app/models/menu"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeMenus(count int) []menu.Menu {

	var objs []menu.Menu

	// 设置唯一性，如 Menu 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		menuModel := menu.Menu{
			// FIXME()
			MenuName:       faker.Name(),
			ParentId:       uint64(rand.Int63n(20)),
			Icon:           faker.Name(),
			Uri:            faker.Name(),
			IsLink:         uint64(rand.Int63n(2)),
			PermissionName: faker.Name(),
			GuardName:      faker.Name(),
			Sequence:       uint64(rand.Int63n(20)),
		}
		objs = append(objs, menuModel)
	}

	return objs
}
