package factories

import (
	"adcenter/app/models/role"

	"github.com/bxcodec/faker/v3"
)

func MakeRoles(count int) []role.Role {

	var objs []role.Role

	// 设置唯一性，如 Role 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		roleModel := role.Role{
			// FIXME()
			RoleName:    faker.Name(),
			GuardName:   faker.Name(),
			Description: faker.Sentence(),
		}
		objs = append(objs, roleModel)
	}

	return objs
}
