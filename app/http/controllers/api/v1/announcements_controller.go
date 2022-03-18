package v1

import (
	"adcenter/app/models/announcement"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementsController struct {
	BaseAPIController
}

func (ctrl *AnnouncementsController) Index(c *gin.Context) {
	announcements := announcement.All()
	response.Data(c, announcements)
}

func (ctrl *AnnouncementsController) Show(c *gin.Context) {
	announcementModel := announcement.Get(c.Param("id"))
	if announcementModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, announcementModel)
}

func (ctrl *AnnouncementsController) Store(c *gin.Context) {

	request := requests.AnnouncementRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementSave); !ok {
		return
	}

	announcementModel := announcement.Announcement{

		AnnouncementNo: request.AnnouncementNo,
		CreatorId:      request.CreatorId,
		DepartmentId:   request.DepartmentId,
		Title:          request.Title,
		LongTitle:      request.LongTitle,
		Type:           request.Type,
		Banner:         request.Banner,
		RedirectTo:     request.RedirectTo,
		RedirectParams: request.RedirectParams,
		Content:        request.Content,
	}
	announcementModel.Create()
	if announcementModel.ID > 0 {
		response.Created(c, announcementModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementsController) Update(c *gin.Context) {

	announcementModel := announcement.Get(c.Param("id"))
	if announcementModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncement(c, announcementModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementSave)
	if !bindOk {
		return
	}

	announcementModel.AnnouncementNo = request.AnnouncementNo
	announcementModel.CreatorId = request.CreatorId
	announcementModel.DepartmentId = request.DepartmentId
	announcementModel.Title = request.Title
	announcementModel.LongTitle = request.LongTitle
	announcementModel.Type = request.Type
	announcementModel.Banner = request.Banner
	announcementModel.RedirectTo = request.RedirectTo
	announcementModel.RedirectParams = request.RedirectParams
	announcementModel.Content = request.Content
	rowsAffected := announcementModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementsController) Delete(c *gin.Context) {

	announcementModel := announcement.Get(c.Param("id"))
	if announcementModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncement(c, announcementModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
