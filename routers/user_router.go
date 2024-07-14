package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

// 自定义密钥
var secret = []byte("hdadhiadadiadhkadh")

func UserRouter(router *gin.RouterGroup) {

	loginApi := api.ApiGroupApp.LoginApi

	// 创建基于cookie的存储引擎，secret 参数是用于加密的密钥，可以随便填写
	store := cookie.NewStore(secret)
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎
	router.Use(sessions.Sessions("mysession", store))

	router.POST("/email_login", loginApi.EmailLoginView)
	router.GET("/users", middleware.JwtAuth(), loginApi.UserListView)
	router.PUT("/user_update_role", middleware.JwtAuth(), loginApi.UserUpdateRoleView)
	router.PUT("/user_update_pwd", middleware.JwtAuth(), loginApi.UserUpdatePwdView)
	router.GET("/user_logout", middleware.JwtAuth(), loginApi.UserLogoutView)
	router.DELETE("/user_delete", middleware.JwtAuth(), loginApi.UserRemoveView)
	router.POST("/user_bind_email", middleware.JwtAuth(), loginApi.UserBindMailView)
	router.POST("/user_register", loginApi.UserRegisterView)

}
