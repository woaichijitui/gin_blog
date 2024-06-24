package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuUpdateView 修改菜单api （只修改数据库的名称）
// @Tags 菜单管理
// @summary 修改菜单
// @Description 修改菜单api
// @Param cr body MenuRequest true "要更新的菜单参数"
// @Param id path string true "要更新的菜单id"
// @Router /menu/{id} [put]
// @Produce json
// @success 200 {object} res.Response
func (MenuAPi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	//		绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//绑定错误 返回msg tag 中错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	//id
	id := c.Param("id")
	var menuModel models.MenuModel
	//	根据id数据库查找该菜单
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMassage("没有该菜单", c)
		return
	}

	//清空menu关联的图片
	err = global.DB.Model(&menuModel).Association("Banners").Clear() //关联的第三章表也一并删除了
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("删除菜单图片失败", c)
		return
	}

	//操作第三张表
	if len(cr.ImageSortList) > 0 {
		list := make([]models.MenuBannerModel, 0)
		for _, sort := range cr.ImageSortList {
			list = append(list, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&list).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMassage("添加第三章表失败", c)
			return
		}
	}

	//更新
	mp := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(&mp).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMassage("修改菜单失败", c)
		return
	}
	//成功响应
	res.OkWithMassage("菜单修改成功", c)

}
