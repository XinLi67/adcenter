package factories

import (
	"adcenter/app/models/permission_group"

	"github.com/bxcodec/faker/v3"
)

func MakePermissionGroups(count int) []permission_group.PermissionGroup {

	var objs []permission_group.PermissionGroup

	// 设置唯一性，如 PermissionGroup 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		permissionGroupModel := permission_group.PermissionGroup{
			PermissionGroupName: faker.Name(),
		}
		objs = append(objs, permissionGroupModel)
	}

	return objs
}
