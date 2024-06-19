package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type AdvertResponse struct {
	Title  string `json:"title" binding:"required" msg:"标题错误" struct:"title"`             //显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转连接非法" struct:"href"`         //跳转连接
	Images string `json:"images" binding:"required,url" msg:"广告图片非法" struct:"images"`     //图片
	IsShow bool   `json:"is_show" binding:"required_bool" msg:"请选择是否展示" struct:"is_show"` //是否展示
}

// AdvertCreateView 创建广告
// @Tags 广告管理
// @summary 创建广告
// @Description 创建广告
// @Param cr body AdvertResponse true "表示多个参数"
// @Router /advert [post]
// @Produce json
// @success 200 {object} res.Response
func (AdvertApi) AdvertCreateView(c *gin.Context) {

	var cr AdvertResponse
	//		绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//绑定错误 返回msg tag 中错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	//	判断重复标题
	err = global.DB.Take(&models.AdvertModel{}, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMassage("标题重复", c)
		return
	}

	//	存入数据库
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		res.FailWithMassage(err.Error(), c)
		return
	}
	//	成功响应
	res.OkWith(c)

}
