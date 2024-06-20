package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImagesRouter(router *gin.RouterGroup) {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.GET("/images", ImagesApi.ImagesListView)
	router.GET("/image_names", ImagesApi.ImagesNameListView)
	router.POST("/images", ImagesApi.ImagesUploadView)
	router.PUT("/images", ImagesApi.ImagesUpdateView)
	router.DELETE("/images", ImagesApi.ImagesRemoveView)
}
