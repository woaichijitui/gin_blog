package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// ImagesRemoveView 批量删除图片接口
func (ImagesApi) ImagesRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	//	绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//	判断文件是否存在
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMassage("文件不存在", c)
		return
	}

	//数据库删除
	global.DB.Delete(imageList)

	//	成功响应
	res.OkWithMassage(fmt.Sprintf("成功删除 %d 张图片", count), c)
}
