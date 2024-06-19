package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertUpdateView 创建广告
// @Tags 广告管理
// @summary 更新广告
// @Description 更新广告
// @Param id path string  true "URL 参数 ：id"
// @Param cr body AdvertResponse  false "创建广告的示例"
// @Router /advert/{id} [put] ""
// @Produce json
// @success 200 {object} res.Response{}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {

	id := c.Param("id")

	var cr AdvertResponse
	//		绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//绑定错误 返回msg tag 中错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	var advert models.AdvertModel
	//根据id查询是否有此广告
	row := global.DB.Take(&advert, "id = ?", id).RowsAffected
	if row == 0 {
		res.OkWithMassage("没有此广告", c)
		return
	}

	//更新
	mp := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(&mp).Error
	if err != nil {
		res.FailWithMassage(err.Error(), c)
		return
	}
	//	成功响应
	res.OkWithMassage("广告更新成功", c)

}
