package v1

import (
	"adcenter/app/models/media"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	BaseAPIController
}

func (ctrl *MediaController) Index(c *gin.Context) {
	media := media.All()
	response.Data(c, media)
}

func (ctrl *MediaController) Show(c *gin.Context) {
	mediaModel := media.Get(c.Param("id"))
	if mediaModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, mediaModel)
}

func (ctrl *MediaController) Store(c *gin.Context) {

	request := requests.MediaRequest{}
	if ok := requests.Validate(c, &request, requests.MediaSave); !ok {
		return
	}

	mediaModel := media.Media{

		CreatorId:    request.CreatorId,
		MediaGroupId: request.MediaGroupId,
		DepartmentId: request.DepartmentId,
		Type:         request.Type,
		Url:          request.Url,
		Title:        request.Title,
		Content:      request.Content,
	}
	mediaModel.Create()
	if mediaModel.ID > 0 {
		response.Created(c, mediaModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MediaController) Update(c *gin.Context) {

	mediaModel := media.Get(c.Param("id"))
	if mediaModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMedia(c, mediaModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.MediaRequest{}
	bindOk := requests.Validate(c, &request, requests.MediaSave)
	if !bindOk {
		return
	}

	mediaModel.CreatorId = request.CreatorId
	mediaModel.MediaGroupId = request.MediaGroupId
	mediaModel.DepartmentId = request.DepartmentId
	mediaModel.Type = request.Type
	mediaModel.Url = request.Url
	mediaModel.Title = request.Title
	mediaModel.Content = request.Content
	rowsAffected := mediaModel.Save()
	if rowsAffected > 0 {
		response.Data(c, mediaModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MediaController) Delete(c *gin.Context) {

	mediaModel := media.Get(c.Param("id"))
	if mediaModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMedia(c, mediaModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := mediaModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
