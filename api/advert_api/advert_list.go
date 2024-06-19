package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/service_com"
	"strings"
)

// AdvertListView 广告list
// @Tags 广告管理
// @summary 广告列表
// @Description 广告列表
// @Param page query models.PageInfo true "表示单个参数"
// @Router /advert [get]
// @Produce json
// @success 200 {object} res.Response{data=[]models.AdvertModel}
func (AdvertApi) AdvertListView(c *gin.Context) {
	//	绑定参数
	var AdvertList []models.AdvertModel
	var page models.PageInfo
	var advert models.AdvertModel
	//绑定参数
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	advert.IsShow = true
	//	不需要展示的
	referer := c.GetHeader("Referer")
	fmt.Println(referer)
	if contain := strings.Contains(referer, "admin"); contain {
		//数据库查不出 advert.IsShow = false 的记录
		advert.IsShow = false

	}
	//	展示
	AdvertList, count, err := service_com.ComList(advert, service_com.Option{
		PageInfo: page,
		Logger:   true,
	})
	//	成功响应
	res.OkWithList(AdvertList, count, c)
}
