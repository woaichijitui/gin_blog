package advert_api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/common"
)

type AdvertResponse struct {
	Title  string `json:"title" binding:"required" msg:"标题错误"`          //显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转连接非法"`     //跳转连接
	Images string `json:"images" binding:"required,url" msg:"广告图片非法"`   //图片
	IsShow bool   `json:"is_show" binding:"requiredbool" msg:"请选择是否展示"` //是否展示
}

// 创建广告
func (AdvertApi) AdvertCreateView(c *gin.Context) {

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("requiredbool", common.RequiredBool)
	}

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
