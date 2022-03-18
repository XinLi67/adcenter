// Package auth 授权相关逻辑
package auth

import (
	"adcenter/app/models/admin_user"
	"adcenter/pkg/logger"
	"errors"

	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(email string, password string) (admin_user.AdminUser, error) {
	adminUserModel := admin_user.GetByMulti(email)
	if adminUserModel.ID == 0 {
		return admin_user.AdminUser{}, errors.New("账号不存在")
	}

	if !adminUserModel.ComparePassword(password) {
		return admin_user.AdminUser{}, errors.New("密码错误")
	}

	return adminUserModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (admin_user.AdminUser, error) {
	adminUserModel := admin_user.GetByPhone(phone)
	if adminUserModel.ID == 0 {
		return admin_user.AdminUser{}, errors.New("手机号未注册")
	}

	return adminUserModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) admin_user.AdminUser {
	adminUserModel, ok := c.MustGet("current_user").(admin_user.AdminUser)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return admin_user.AdminUser{}
	}
	// db is now a *DB value
	return adminUserModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
