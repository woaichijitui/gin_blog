package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type BannerResponse struct {
	ID   uint   `json:"id"`   //图片id
	Path string `json:"path"` //图片路径
	Name string `json:"name"` //图片名称
}

// ImagesNameListView 图片名字列表（全部）查询接口
// @Tags 图片管理
// @summary 图片名字列表
// @Description 图片名字列表
// @Router /image_names [get]
// @Produce json
// @success 200 {object} res.Response{data=[]BannerResponse}
func (ImagesApi) ImagesNameListView(c *gin.Context) {

	//响应的图片列表
	var imageList []BannerResponse
	//数据库查看全部的图片列表
	//Scan方法将所有选择的记录扫描到结构体
	global.DB.Model(&models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)

	//	成功响应
	res.OkWithData(imageList, c)

}
