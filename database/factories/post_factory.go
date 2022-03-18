package factories

import (
	"adcenter/app/models/post"

	"github.com/bxcodec/faker/v3"
)

func MakePosts(count int) []post.Post {

	var objs []post.Post

	// 设置唯一性，如 Post 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		postModel := post.Post{
			//这里需要添加内容
			Title:   faker.Name(),
			Content: faker.Paragraph(),
		}
		objs = append(objs, postModel)
	}

	return objs
}
