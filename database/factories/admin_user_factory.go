package factories

import (
	"adcenter/app/models/admin_user"
	"adcenter/pkg/helpers"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeAdminUsers(count int) []admin_user.AdminUser {

	var objs []admin_user.AdminUser

	// 设置唯一性，如 AdminUser 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		adminUserModel := admin_user.AdminUser{
			// FIXME()
			Name:         faker.Name(),
			Username:     faker.Username(),
			DepartmentId: uint64(rand.Int63n(10)),
			Gender:       uint64(rand.Int63n(2)),
			Email:        faker.Email(),
			Phone:        helpers.RandomNumber(11),
			Password:     faker.Password(),
			Status:       uint64(rand.Int63n(2)),
		}
		objs = append(objs, adminUserModel)
	}

	return objs
}
