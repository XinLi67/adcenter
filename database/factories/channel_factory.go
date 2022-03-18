package factories

import (
	"adcenter/app/models/channel"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func MakeChannels(count int) []channel.Channel {

	var objs []channel.Channel

	// 设置唯一性，如 Channel 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		channelModel := channel.Channel{
			// FIXME()
			ChannelName: faker.Name(),
			Status:      uint64(rand.Int63n(2)),
			Description: faker.Sentence(),
		}
		objs = append(objs, channelModel)
	}

	return objs
}
