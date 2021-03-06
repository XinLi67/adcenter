package v1

import (
	"adcenter/app/models/click_record"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type ClickRecordsController struct {
	BaseAPIController
}

func (ctrl *ClickRecordsController) Index(c *gin.Context) {
	clickRecords := click_record.All()
	response.Data(c, clickRecords)
}

func (ctrl *ClickRecordsController) Show(c *gin.Context) {
	clickRecordModel := click_record.Get(c.Param("id"))
	if clickRecordModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, clickRecordModel)
}

func (ctrl *ClickRecordsController) Store(c *gin.Context) {

	request := requests.ClickRecordRequest{}
	if ok := requests.Validate(c, &request, requests.ClickRecordSave); !ok {
		return
	}

	clickRecordModel := click_record.ClickRecord{

		AdvertisingId: request.AdvertisingId,
		CustomerId:    request.CustomerId,
		BrowsingTime:  request.BrowsingTime,
		StartTime:     request.StartTime,
		EndTime:       request.EndTime,
	}
	clickRecordModel.Create()
	if clickRecordModel.ID > 0 {
		response.Created(c, clickRecordModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ClickRecordsController) Update(c *gin.Context) {

	clickRecordModel := click_record.Get(c.Param("id"))
	if clickRecordModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyClickRecord(c, clickRecordModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.ClickRecordRequest{}
	bindOk := requests.Validate(c, &request, requests.ClickRecordSave)
	if !bindOk {
		return
	}

	clickRecordModel.AdvertisingId = request.AdvertisingId
	clickRecordModel.CustomerId = request.CustomerId
	clickRecordModel.BrowsingTime = request.BrowsingTime
	clickRecordModel.StartTime = request.StartTime
	clickRecordModel.EndTime = request.EndTime
	rowsAffected := clickRecordModel.Save()
	if rowsAffected > 0 {
		response.Data(c, clickRecordModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ClickRecordsController) Delete(c *gin.Context) {

	clickRecordModel := click_record.Get(c.Param("id"))
	if clickRecordModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyClickRecord(c, clickRecordModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := clickRecordModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
