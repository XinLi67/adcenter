package v1

import (
	"adcenter/app/models/menu"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type MenusController struct {
	BaseAPIController
}

func (ctrl *MenusController) Index(c *gin.Context) {
	menus := menu.All()
	response.Data(c, menus)
}

func (ctrl *MenusController) Show(c *gin.Context) {
	menuModel := menu.Get(c.Param("id"))
	if menuModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, menuModel)
}

func (ctrl *MenusController) Store(c *gin.Context) {

	request := requests.MenuRequest{}
	if ok := requests.Validate(c, &request, requests.MenuSave); !ok {
		return
	}

	menuModel := menu.Menu{

		MenuName:       request.MenuName,
		ParentId:       request.ParentId,
		Icon:           request.Icon,
		Uri:            request.Uri,
		IsLink:         request.IsLink,
		PermissionName: request.PermissionName,
		GuardName:      request.GuardName,
		Sequence:       request.Sequence,
	}
	menuModel.Create()
	if menuModel.ID > 0 {
		response.Created(c, menuModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MenusController) Update(c *gin.Context) {

	menuModel := menu.Get(c.Param("id"))
	if menuModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMenu(c, menuModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.MenuRequest{}
	bindOk := requests.Validate(c, &request, requests.MenuSave)
	if !bindOk {
		return
	}

	menuModel.MenuName = request.MenuName
	menuModel.ParentId = request.ParentId
	menuModel.Icon = request.Icon
	menuModel.Uri = request.Uri
	menuModel.IsLink = request.IsLink
	menuModel.PermissionName = request.PermissionName
	menuModel.GuardName = request.GuardName
	menuModel.Sequence = request.Sequence
	rowsAffected := menuModel.Save()
	if rowsAffected > 0 {
		response.Data(c, menuModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MenusController) Delete(c *gin.Context) {

	menuModel := menu.Get(c.Param("id"))
	if menuModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMenu(c, menuModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := menuModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
