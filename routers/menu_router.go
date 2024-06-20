package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MenuRouter(router *gin.RouterGroup) {

	menuAPi := api.ApiGroupApp.MenuAPi
	router.GET("/menu", menuAPi.MenuListView)
	router.POST("/menu", menuAPi.MenuCreateView)

}
