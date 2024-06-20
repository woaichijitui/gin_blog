package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// UpdateNameResponse
// Description: 修改图片名称的请求结构体
type UpdateNameResponse struct {
	Id   uint   `json:"id" binding:"required" msg:"请选择文件的id"`
	Name string `json:"name" binding:"required,max=256" msg:"请输入文件名，或名字太长"`
}

// ImagesUpdateView 修改图片名字api （只修改数据库的名称）
// @Tags 图片管理
// @summary 修改图片
// @Description 修改图片
// @Param cr body UpdateNameResponse true "要更新的图片id和name"
// @Router /images/{id} [put]
// @Produce json
// @success 200 {object} res.Response
func (ImagesApi) ImagesUpdateView(c *gin.Context) {
	var cr UpdateNameResponse
	//		绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//绑定错误 返回msg tag 中错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	var banner models.BannerModel
	//	数据库查找该文件
	err = global.DB.Take(&banner, cr.Id).Error
	if err != nil {
		res.FailWithMassage("没有该文件", c)
		return
	}

	//	修改名称
	//	问题：若修改后的名称没有带.jpg等图片格式后缀，会对后续影响吗？
	err = global.DB.Model(&banner).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMassage(err.Error(), c)
		return
	}
	//成功响应
	res.OkWithMassage("图片名字修改成功", c)

}
