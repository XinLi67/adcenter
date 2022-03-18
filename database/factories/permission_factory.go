package factories

import (
	"adcenter/app/models/permission"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakePermissions(count int) []permission.Permission {

	var objs []permission.Permission

	// 设置唯一性，如 Permission 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		permissionModel := permission.Permission{
			// FIXME()
			PermissionGroupId: uint64(rand.Int63n(10)),
			PermissionName:    faker.Name(),
			Icon:              faker.Name(),
			GuardName:         faker.Name(),
			DisplayName:       faker.Name(),
			Sequence:          uint64(rand.Int63n(20)),
			Description:       faker.Sentence(),
			CreatedName:       faker.FirstName(),
			UpdatedName:       faker.FirstName(),
		}
		objs = append(objs, permissionModel)
	}

	return objs
}
