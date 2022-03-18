package v1

import (
	"adcenter/app/models/announcement_plan"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementPlansController struct {
	BaseAPIController
}

func (ctrl *AnnouncementPlansController) Index(c *gin.Context) {
	announcementPlans := announcement_plan.All()
	response.Data(c, announcementPlans)
}

func (ctrl *AnnouncementPlansController) Show(c *gin.Context) {
	announcementPlanModel := announcement_plan.Get(c.Param("id"))
	if announcementPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, announcementPlanModel)
}

func (ctrl *AnnouncementPlansController) Store(c *gin.Context) {

	request := requests.AnnouncementPlanRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementPlanSave); !ok {
		return
	}

	announcementPlanModel := announcement_plan.AnnouncementPlan{

		CreatorId:              request.CreatorId,
		AnnouncementId:         request.AnnouncementId,
		AnnouncementType:       request.AnnouncementType,
		AnnouncementPositionId: request.AnnouncementPositionId,
		Order:                  request.Order,
		SchedulingDate:         request.SchedulingDate,
		SchedulingTime:         request.SchedulingTime,
		StartDate:              request.StartDate,
		EndTDate:               request.EndTDate,
		StartTime:              request.StartTime,
		EndTime:                request.EndTime,
		Status:                 request.Status,
	}
	announcementPlanModel.Create()
	if announcementPlanModel.ID > 0 {
		response.Created(c, announcementPlanModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPlansController) Update(c *gin.Context) {

	announcementPlanModel := announcement_plan.Get(c.Param("id"))
	if announcementPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPlan(c, announcementPlanModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementPlanRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementPlanSave)
	if !bindOk {
		return
	}

	announcementPlanModel.CreatorId = request.CreatorId
	announcementPlanModel.AnnouncementId = request.AnnouncementId
	announcementPlanModel.AnnouncementType = request.AnnouncementType
	announcementPlanModel.AnnouncementPositionId = request.AnnouncementPositionId
	announcementPlanModel.Order = request.Order
	announcementPlanModel.SchedulingDate = request.SchedulingDate
	announcementPlanModel.SchedulingTime = request.SchedulingTime
	announcementPlanModel.StartDate = request.StartDate
	announcementPlanModel.EndTDate = request.EndTDate
	announcementPlanModel.StartTime = request.StartTime
	announcementPlanModel.EndTime = request.EndTime
	announcementPlanModel.Status = request.Status
	rowsAffected := announcementPlanModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementPlanModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPlansController) Delete(c *gin.Context) {

	announcementPlanModel := announcement_plan.Get(c.Param("id"))
	if announcementPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPlan(c, announcementPlanModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPlanModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
