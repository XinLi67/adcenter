package v1

import (
	"adcenter/app/models/permission"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionsController struct {
	BaseAPIController
}

func (ctrl *PermissionsController) Index(c *gin.Context) {
	permissions := permission.All()
	response.Data(c, permissions)
}

func (ctrl *PermissionsController) Show(c *gin.Context) {
	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, permissionModel)
}

func (ctrl *PermissionsController) Store(c *gin.Context) {

	request := requests.PermissionRequest{}
	if ok := requests.Validate(c, &request, requests.PermissionSave); !ok {
		return
	}

	permissionModel := permission.Permission{

		PermissionGroupId: request.PermissionGroupId,
		PermissionName:    request.PermissionName,
		Icon:              request.Icon,
		GuardName:         request.GuardName,
		DisplayName:       request.DisplayName,
		Sequence:          request.Sequence,
		Description:       request.Description,
		CreatedName:       request.CreatedName,
		UpdatedName:       request.UpdatedName,
	}
	permissionModel.Create()
	if permissionModel.ID > 0 {
		response.Created(c, permissionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Update(c *gin.Context) {

	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermission(c, permissionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.PermissionRequest{}
	bindOk := requests.Validate(c, &request, requests.PermissionSave)
	if !bindOk {
		return
	}

	permissionModel.PermissionGroupId = request.PermissionGroupId
	permissionModel.PermissionName = request.PermissionName
	permissionModel.Icon = request.Icon
	permissionModel.GuardName = request.GuardName
	permissionModel.DisplayName = request.DisplayName
	permissionModel.Sequence = request.Sequence
	permissionModel.Description = request.Description
	permissionModel.CreatedName = request.CreatedName
	permissionModel.UpdatedName = request.UpdatedName
	rowsAffected := permissionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, permissionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Delete(c *gin.Context) {

	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermission(c, permissionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := permissionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
