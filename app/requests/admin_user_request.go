package requests

import (
	"adcenter/app/requests/validators"
	"adcenter/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdminUserRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	// FIXME()
	Name         string `valid:"name" json:"name,omitempty"`
	Username     string `valid:"user_name" json:"user_name,omitempty"`
	DepartmentId uint64 `valid:"department_id" json:"department_id,omitempty"`
	Gender       uint64 `valid:"gender" json:"gender,omitempty"`
	Email        string `valid:"email" json:"email,omitempty"`
	Phone        string `valid:"phone" json:"phone,omitempty"`
	Password     string `valid:"password" json:"password,omitempty"`
	Status       uint64 `valid:"status" json:"status,omitempty"`
	// VerifyCode   string `valid:"verify_code" json:"verify_code,omitempty"`
}

func AdminUserSave(data interface{}, c *gin.Context) map[string][]string {

	// 查询用户名重复时，过滤掉当前用户 ID
	uid := auth.CurrentUID(c)
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"user_name": []string{
			"required",
			"alpha_num",
			"between:3,20",
			"not_exists:admin_users,name," + uid,
		},
		"email": []string{
			"email",
			"not_exists:admin_users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		// "verify_code": []string{
		// 	"required",
		// 	"digits:6",
		// },
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:admin_users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
	}
	messages := govalidator.MapData{
		"user_name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已被占用",
		},
		"email": []string{
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
			"not_in:新的 Email 与老 Email 一致",
		},
		// "verify_code": []string{
		// 	"required:验证码答案必填",
		// 	"digits:验证码长度必须为 6 位的数字",
		// },
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机与老手机号一致",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdatePasswordRequest struct {
	Password           string `valid:"password" json:"password,omitempty"`
	NewPassword        string `valid:"new_password" json:"new_password,omitempty"`
	NewPasswordConfirm string `valid:"new_password_confirm" json:"new_password_confirm,omitempty"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password_confirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 6",
		},
	}

	// 确保 comfirm 密码正确
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)

	return errs
}
