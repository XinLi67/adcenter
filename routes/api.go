// Package routes 注册路由
package routes

import (
	controllers "adcenter/app/http/controllers/api/v1"
	"adcenter/app/http/controllers/api/v1/auth"
	"adcenter/app/http/middlewares"
	"adcenter/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)
	// post := new(controllers.PostsController)

	// 查询API
	v1.GET("/posts/:id", post.Show)
	v1.GET("/admin-user/:id", post.Show)
	v1.GET("/advertising/:id", post.Show)
	v1.GET("/announcement/:id", post.Show)
	v1.GET("/advertising-position/:id", post.Show)
	v1.GET("/announcement-position/:id", post.Show)
	v1.GET("/advertising-plan/:id", post.Show)
	v1.GET("/announcement-plan/:id", post.Show)
	v1.GET("/channel/:id", post.Show)
	v1.GET("/media/:id", post.Show)
	v1.GET("/media-group/:id", post.Show)
	v1.GET("/click-record/:id", post.Show)
	v1.GET("/audit-record/:id", post.Show)
	v1.GET("/department/:id", post.Show)
	v1.GET("/role/:id", post.Show)
	v1.GET("/permission/:id", post.Show)
	v1.GET("/permission-group/:id", post.Show)
	v1.GET("/menu/:id", post.Show)

	v1.POST("/posts", post.Store)
	v1.POST("/posts/:id/update", post.Update)
	v1.POST("/posts/:id/delete", post.Delete)

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)

			// 注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
		}

		uc := new(controllers.UsersController)

		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewares.AuthJWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJWT(), uc.UpdatePassword)
			// usersGroup.PUT("/avatar", middlewares.AuthJWT(), uc.UpdateAvatar)
		}
	}
}
