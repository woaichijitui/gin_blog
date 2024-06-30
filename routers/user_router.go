package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func LoginRouter(router *gin.RouterGroup) {

	loginApi := api.ApiGroupApp.LoginApi
	router.POST("/email_login", loginApi.EmailLoginView)
	router.GET("/users", middleware.JwtAdmin(), loginApi.UserListView)
	router.PUT("/user_update_role", middleware.JwtAdmin(), loginApi.UserUpdateRole)
	router.GET("/user_logout", middleware.JwtAdmin(), loginApi.UserLogoutView)

}
