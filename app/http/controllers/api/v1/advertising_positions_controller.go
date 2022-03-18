package v1

import (
	"adcenter/app/models/advertising_position"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdvertisingPositionsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingPositionsController) Index(c *gin.Context) {
	advertisingPositions := advertising_position.All()
	response.Data(c, advertisingPositions)
}

func (ctrl *AdvertisingPositionsController) Show(c *gin.Context) {
	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, advertisingPositionModel)
}

func (ctrl *AdvertisingPositionsController) Store(c *gin.Context) {

	request := requests.AdvertisingPositionRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingPositionSave); !ok {
		return
	}

	advertisingPositionModel := advertising_position.AdvertisingPosition{

		AdvertisingPositionName: request.AdvertisingPositionName,
		ParentId:                request.ParentId,
		Code:                    request.Code,
		Height:                  request.Height,
		Weight:                  request.Weight,
	}
	advertisingPositionModel.Create()
	if advertisingPositionModel.ID > 0 {
		response.Created(c, advertisingPositionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPositionsController) Update(c *gin.Context) {

	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingPositionRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingPositionSave)
	if !bindOk {
		return
	}

	advertisingPositionModel.AdvertisingPositionName = request.AdvertisingPositionName
	advertisingPositionModel.ParentId = request.ParentId
	advertisingPositionModel.Code = request.Code
	advertisingPositionModel.Height = request.Height
	advertisingPositionModel.Weight = request.Weight

	rowsAffected := advertisingPositionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingPositionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPositionsController) Delete(c *gin.Context) {

	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPositionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
