package v1

import (
	"adcenter/app/models/advertising_plan"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdvertisingPlansController struct {
	BaseAPIController
}

func (ctrl *AdvertisingPlansController) Index(c *gin.Context) {
	advertisingPlans := advertising_plan.All()
	response.Data(c, advertisingPlans)
}

func (ctrl *AdvertisingPlansController) Show(c *gin.Context) {
	advertisingPlanModel := advertising_plan.Get(c.Param("id"))
	if advertisingPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, advertisingPlanModel)
}

func (ctrl *AdvertisingPlansController) Store(c *gin.Context) {

	request := requests.AdvertisingPlanRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingPlanSave); !ok {
		return
	}

	advertisingPlanModel := advertising_plan.AdvertisingPlan{

		CreatorId:             request.CreatorId,
		AdvertisingId:         request.AdvertisingId,
		AdvertisingType:       request.AdvertisingType,
		AdvertisingPositionId: request.AdvertisingPositionId,
		Order:                 request.Order,
		SchedulingDate:        request.SchedulingDate,
		SchedulingTime:        request.SchedulingTime,
		StartDate:             request.StartDate,
		EndTDate:              request.EndTDate,
		StartTime:             request.StartTime,
		EndTime:               request.EndTime,
		Status:                request.Status,
	}
	advertisingPlanModel.Create()
	if advertisingPlanModel.ID > 0 {
		response.Created(c, advertisingPlanModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPlansController) Update(c *gin.Context) {

	advertisingPlanModel := advertising_plan.Get(c.Param("id"))
	if advertisingPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPlan(c, advertisingPlanModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingPlanRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingPlanSave)
	if !bindOk {
		return
	}

	advertisingPlanModel.CreatorId = request.CreatorId
	advertisingPlanModel.AdvertisingId = request.AdvertisingId
	advertisingPlanModel.AdvertisingType = request.AdvertisingType
	advertisingPlanModel.AdvertisingPositionId = request.AdvertisingPositionId
	advertisingPlanModel.Order = request.Order
	advertisingPlanModel.SchedulingDate = request.SchedulingDate
	advertisingPlanModel.SchedulingTime = request.SchedulingTime
	advertisingPlanModel.StartDate = request.StartDate
	advertisingPlanModel.EndTDate = request.EndTDate
	advertisingPlanModel.StartTime = request.StartTime
	advertisingPlanModel.EndTime = request.EndTime
	advertisingPlanModel.Status = request.Status
	rowsAffected := advertisingPlanModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingPlanModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPlansController) Delete(c *gin.Context) {

	advertisingPlanModel := advertising_plan.Get(c.Param("id"))
	if advertisingPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPlan(c, advertisingPlanModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPlanModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
