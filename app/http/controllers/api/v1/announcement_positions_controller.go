package v1

import (
	"adcenter/app/models/announcement_position"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementPositionsController struct {
	BaseAPIController
}

func (ctrl *AnnouncementPositionsController) Index(c *gin.Context) {
	announcementPositions := announcement_position.All()
	response.Data(c, announcementPositions)
}

func (ctrl *AnnouncementPositionsController) Show(c *gin.Context) {
	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, announcementPositionModel)
}

func (ctrl *AnnouncementPositionsController) Store(c *gin.Context) {

	request := requests.AnnouncementPositionRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementPositionSave); !ok {
		return
	}

	announcementPositionModel := announcement_position.AnnouncementPosition{

		AnnouncementPositionName: request.AnnouncementPositionName,
		ParentId:                 request.ParentId,
		Code:                     request.Code,
		Height:                   request.Height,
		Weight:                   request.Weight,
	}
	announcementPositionModel.Create()
	if announcementPositionModel.ID > 0 {
		response.Created(c, announcementPositionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPositionsController) Update(c *gin.Context) {

	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementPositionRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementPositionSave)
	if !bindOk {
		return
	}

	announcementPositionModel.AnnouncementPositionName = request.AnnouncementPositionName
	announcementPositionModel.ParentId = request.ParentId
	announcementPositionModel.Code = request.Code
	announcementPositionModel.Height = request.Height
	announcementPositionModel.Weight = request.Weight
	rowsAffected := announcementPositionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementPositionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPositionsController) Delete(c *gin.Context) {

	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPositionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
