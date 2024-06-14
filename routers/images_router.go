package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImagesRouter(router *gin.RouterGroup) {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images", ImagesApi.ImagesUploadView)
}
