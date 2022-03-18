package v1

import (
	"adcenter/app/models/audit_record"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuditRecordsController struct {
	BaseAPIController
}

func (ctrl *AuditRecordsController) Index(c *gin.Context) {
	auditRecords := audit_record.All()
	response.Data(c, auditRecords)
}

func (ctrl *AuditRecordsController) Show(c *gin.Context) {
	auditRecordModel := audit_record.Get(c.Param("id"))
	if auditRecordModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, auditRecordModel)
}

func (ctrl *AuditRecordsController) Store(c *gin.Context) {

	request := requests.AuditRecordRequest{}
	if ok := requests.Validate(c, &request, requests.AuditRecordSave); !ok {
		return
	}

	auditRecordModel := audit_record.AuditRecord{

		AuditableId:   request.AuditableId,
		AuditableType: request.AuditableType,
		ApplicantId:   request.ApplicantId,
		AuditorId:     request.AuditorId,
		Status:        request.Status,
		Content:       request.Content,
	}
	auditRecordModel.Create()
	if auditRecordModel.ID > 0 {
		response.Created(c, auditRecordModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AuditRecordsController) Update(c *gin.Context) {

	auditRecordModel := audit_record.Get(c.Param("id"))
	if auditRecordModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAuditRecord(c, auditRecordModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AuditRecordRequest{}
	bindOk := requests.Validate(c, &request, requests.AuditRecordSave)
	if !bindOk {
		return
	}

	auditRecordModel.AuditableId = request.AuditableId
	auditRecordModel.AuditableType = request.AuditableType
	auditRecordModel.ApplicantId = request.ApplicantId
	auditRecordModel.AuditorId = request.AuditorId
	auditRecordModel.Status = request.Status
	auditRecordModel.Content = request.Content
	rowsAffected := auditRecordModel.Save()
	if rowsAffected > 0 {
		response.Data(c, auditRecordModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AuditRecordsController) Delete(c *gin.Context) {

	auditRecordModel := audit_record.Get(c.Param("id"))
	if auditRecordModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAuditRecord(c, auditRecordModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := auditRecordModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
