package factories

import (
	"adcenter/app/models/advertising"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeAdvertisings(count int) []advertising.Advertising {

	var objs []advertising.Advertising

	// 设置唯一性，如 Advertising 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		advertisingModel := advertising.Advertising{
			//FIXME()
			AdvertisingNo:  uint64(rand.Int63n(200)),
			CreatorId:      uint64(rand.Int63n(10)), //10表示在区间[0,10)中生成一个随机整数，前开后闭，不含10
			DepartmentId:   uint64(rand.Int63n(10)),
			Title:          faker.Name(),
			Type:           uint64(rand.Int63n(4)), //由于区间为前开后闭，将3改为4
			RedirectTo:     uint64(rand.Int63n(2)), //由于区间为前开后闭，将1改为2
			MediaId:        uint64(rand.Int63n(10)),
			MediaType:      uint64(rand.Int63n(2)), //由于区间为前开后闭，将1改为2
			Size:           faker.Name(),
			RedirectParams: faker.Name(),
			Description:    faker.Paragraph(), //为varchar类型时，不能使用Paragraph方法，否则长度不够；若要使用，则需改为text类型
		}
		objs = append(objs, advertisingModel)
	}

	return objs
}
