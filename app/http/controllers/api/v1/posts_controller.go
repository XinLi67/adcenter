package v1

import (
	"adcenter/app/models/post"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type PostsController struct {
	BaseAPIController
}

func (ctrl *PostsController) Index(c *gin.Context) {
	posts := post.All()
	response.Data(c, posts)
}

func (ctrl *PostsController) Show(c *gin.Context) {
	postModel := post.Get(c.Param("id"))
	if postModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, postModel)
}

func (ctrl *PostsController) Store(c *gin.Context) {

	request := requests.PostRequest{}
	if ok := requests.Validate(c, &request, requests.PostSave); !ok {
		return
	}

	postModel := post.Post{
		Title:   request.Title,
		Content: request.Content,
	}
	postModel.Create()
	if postModel.ID > 0 {
		response.Created(c, postModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PostsController) Update(c *gin.Context) {

	postModel := post.Get(c.Param("id"))
	if postModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyPost(c, postModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := requests.PostRequest{}
	bindOk := requests.Validate(c, &request, requests.PostSave)
	if !bindOk {
		return
	}

	postModel.Title = request.Title
	postModel.Content = request.Content
	rowsAffected := postModel.Save()
	if rowsAffected > 0 {
		response.Data(c, postModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PostsController) Delete(c *gin.Context) {

	postModel := post.Get(c.Param("id"))
	if postModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyPost(c, postModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	rowsAffected := postModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
