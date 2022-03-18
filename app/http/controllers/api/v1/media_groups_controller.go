package v1

import (
	"adcenter/app/models/media_group"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type MediaGroupsController struct {
	BaseAPIController
}

func (ctrl *MediaGroupsController) Index(c *gin.Context) {
	mediaGroups := media_group.All()
	response.Data(c, mediaGroups)
}

func (ctrl *MediaGroupsController) Show(c *gin.Context) {
	mediaGroupModel := media_group.Get(c.Param("id"))
	if mediaGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, mediaGroupModel)
}

func (ctrl *MediaGroupsController) Store(c *gin.Context) {

	request := requests.MediaGroupRequest{}
	if ok := requests.Validate(c, &request, requests.MediaGroupSave); !ok {
		return
	}

	mediaGroupModel := media_group.MediaGroup{

		MediaGroupName: request.MediaGroupName,
	}
	mediaGroupModel.Create()
	if mediaGroupModel.ID > 0 {
		response.Created(c, mediaGroupModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MediaGroupsController) Update(c *gin.Context) {

	mediaGroupModel := media_group.Get(c.Param("id"))
	if mediaGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMediaGroup(c, mediaGroupModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.MediaGroupRequest{}
	bindOk := requests.Validate(c, &request, requests.MediaGroupSave)
	if !bindOk {
		return
	}

	mediaGroupModel.MediaGroupName = request.MediaGroupName
	rowsAffected := mediaGroupModel.Save()
	if rowsAffected > 0 {
		response.Data(c, mediaGroupModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MediaGroupsController) Delete(c *gin.Context) {

	mediaGroupModel := media_group.Get(c.Param("id"))
	if mediaGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMediaGroup(c, mediaGroupModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := mediaGroupModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
