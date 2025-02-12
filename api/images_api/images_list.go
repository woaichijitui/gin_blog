package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/service_com"
)

// ImagesListView 图片列表查询接口
// @Tags 图片管理
// @summary 图片列表
// @Description 图片列表
// @Param page query models.PageInfo true "图片参数"
// @Router /images [get]
// @Produce json
// @success 200 {object} res.Response{data=[]models.BannerModel}
func (ImagesApi) ImagesListView(c *gin.Context) {
	var imagesList []models.BannerModel
	var page models.PageInfo
	//绑定参数
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 查询分页数据
	imagesList, count, _ := service_com.ComList(models.BannerModel{}, service_com.Option{
		PageInfo: page,
		Logger:   true,
	})

	//	成功响应
	res.OkWithList(imagesList, count, c)

}
