package v1

import (
	"adcenter/app/models/advertising"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdvertisingsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingsController) Index(c *gin.Context) {
	advertisings := advertising.All()
	response.Data(c, advertisings)
}

func (ctrl *AdvertisingsController) Show(c *gin.Context) {
	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, advertisingModel)
}

func (ctrl *AdvertisingsController) Store(c *gin.Context) {

	request := requests.AdvertisingRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingSave); !ok {
		return
	}

	advertisingModel := advertising.Advertising{
		AdvertisingNo:  request.AdvertisingNo,
		CreatorId:      request.CreatorId,
		DepartmentId:   request.DepartmentId,
		Title:          request.Title,
		Type:           request.Type,
		RedirectTo:     request.RedirectTo,
		MediaId:        request.MediaId,
		MediaType:      request.MediaType,
		Size:           request.Size,
		RedirectParams: request.RedirectParams,
		Description:    request.Description,
	}
	advertisingModel.Create()
	if advertisingModel.ID > 0 {
		response.Created(c, advertisingModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingsController) Update(c *gin.Context) {

	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingSave)
	if !bindOk {
		return
	}

	advertisingModel.AdvertisingNo = request.AdvertisingNo
	advertisingModel.CreatorId = request.CreatorId
	advertisingModel.DepartmentId = request.DepartmentId
	advertisingModel.Title = request.Title
	advertisingModel.Type = request.Type
	advertisingModel.RedirectTo = request.RedirectTo
	advertisingModel.MediaId = request.MediaId
	advertisingModel.MediaType = request.MediaType
	advertisingModel.Size = request.Size
	advertisingModel.RedirectParams = request.RedirectParams
	advertisingModel.Description = request.Description
	rowsAffected := advertisingModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingsController) Delete(c *gin.Context) {

	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
