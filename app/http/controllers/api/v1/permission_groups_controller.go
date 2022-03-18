package v1

import (
	"adcenter/app/models/permission_group"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionGroupsController struct {
	BaseAPIController
}

func (ctrl *PermissionGroupsController) Index(c *gin.Context) {
	permissionGroups := permission_group.All()
	response.Data(c, permissionGroups)
}

func (ctrl *PermissionGroupsController) Show(c *gin.Context) {
	permissionGroupModel := permission_group.Get(c.Param("id"))
	if permissionGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, permissionGroupModel)
}

func (ctrl *PermissionGroupsController) Store(c *gin.Context) {

	request := requests.PermissionGroupRequest{}
	if ok := requests.Validate(c, &request, requests.PermissionGroupSave); !ok {
		return
	}

	permissionGroupModel := permission_group.PermissionGroup{

		PermissionGroupName: request.PermissionGroupName,
	}
	permissionGroupModel.Create()
	if permissionGroupModel.ID > 0 {
		response.Created(c, permissionGroupModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PermissionGroupsController) Update(c *gin.Context) {

	permissionGroupModel := permission_group.Get(c.Param("id"))
	if permissionGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermissionGroup(c, permissionGroupModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.PermissionGroupRequest{}
	bindOk := requests.Validate(c, &request, requests.PermissionGroupSave)
	if !bindOk {
		return
	}

	permissionGroupModel.PermissionGroupName = request.PermissionGroupName
	rowsAffected := permissionGroupModel.Save()
	if rowsAffected > 0 {
		response.Data(c, permissionGroupModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionGroupsController) Delete(c *gin.Context) {

	permissionGroupModel := permission_group.Get(c.Param("id"))
	if permissionGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermissionGroup(c, permissionGroupModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := permissionGroupModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
