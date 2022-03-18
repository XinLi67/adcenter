package v1

import (
	"adcenter/app/models/admin_user"
	"adcenter/app/policies"
	"adcenter/app/requests"
	"adcenter/pkg/auth"
	"adcenter/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminUsersController struct {
	BaseAPIController
}

func (ctrl *AdminUsersController) Index(c *gin.Context) {
	adminUsers := admin_user.All()
	response.Data(c, adminUsers)
}

func (ctrl *AdminUsersController) Show(c *gin.Context) {
	adminUserModel := admin_user.Get(c.Param("id"))
	if adminUserModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, adminUserModel)
}

func (ctrl *AdminUsersController) Store(c *gin.Context) {

	request := requests.AdminUserRequest{}
	if ok := requests.Validate(c, &request, requests.AdminUserSave); !ok {
		return
	}

	adminUserModel := admin_user.AdminUser{
		Name:         request.Name,
		Username:     request.Username,
		DepartmentId: request.DepartmentId,
		Gender:       request.Gender,
		Email:        request.Email,
		Phone:        request.Phone,
		Password:     request.Password,
		Status:       request.Status,
	}
	adminUserModel.Create()
	if adminUserModel.ID > 0 {
		response.Created(c, adminUserModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdminUsersController) Update(c *gin.Context) {

	adminUserModel := admin_user.Get(c.Param("id"))
	if adminUserModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdminUser(c, adminUserModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdminUserRequest{}
	// 以下这句只有一个返回值
	// bindOk, errs := requests.Validate(c, &request, requests.AdminUserSave)
	bindOk := requests.Validate(c, &request, requests.AdminUserSave)
	if !bindOk {
		return
	}

	adminUserModel.Name = request.Name
	adminUserModel.Username = request.Username
	adminUserModel.DepartmentId = request.DepartmentId
	adminUserModel.Gender = request.Gender
	adminUserModel.Email = request.Email
	adminUserModel.Phone = request.Phone
	adminUserModel.Password = request.Password
	adminUserModel.Status = request.Status
	rowsAffected := adminUserModel.Save()
	if rowsAffected > 0 {
		response.Data(c, adminUserModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdminUsersController) Delete(c *gin.Context) {

	adminUserModel := admin_user.Get(c.Param("id"))
	if adminUserModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdminUser(c, adminUserModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := adminUserModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {

	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}
