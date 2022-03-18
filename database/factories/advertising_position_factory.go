package factories

import (
	"adcenter/app/models/advertising_position"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeAdvertisingPositions(count int) []advertising_position.AdvertisingPosition {

	var objs []advertising_position.AdvertisingPosition

	// 设置唯一性，如 AdvertisingPosition 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		advertisingPositionModel := advertising_position.AdvertisingPosition{
			// FIXME()
			AdvertisingPositionName: faker.Name(),
			ParentId:                uint64(rand.Int63n(20)),
			Code:                    faker.Name(),
			Height:                  uint64(rand.Int63n(1000)),
			Weight:                  uint64(rand.Int63n(1000)),
		}
		objs = append(objs, advertisingPositionModel)
	}

	return objs
}
