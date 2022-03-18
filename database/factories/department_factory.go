package factories

import (
	"adcenter/app/models/department"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeDepartments(count int) []department.Department {

	var objs []department.Department

	// 设置唯一性，如 Department 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		departmentModel := department.Department{
			// FIXME()
			DepartmentName: faker.Name(),
			ParentId:       uint64(rand.Int63n(20)),
		}
		objs = append(objs, departmentModel)
	}

	return objs
}
