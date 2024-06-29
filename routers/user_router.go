package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func LoginRouter(router *gin.RouterGroup) {

	loginApi := api.ApiGroupApp.LoginApi
	router.POST("/email_login", loginApi.EmailLoginView)
	router.GET("/users", loginApi.UserListView)

}
