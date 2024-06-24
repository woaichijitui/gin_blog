package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

// ImageSort 菜单图片的排序(每个菜单下面图片会安时间切换
type ImageSort struct {
	ImageID uint `json:"image_id" structs:"image_id"` //图片id
	Sort    int  `json:"sort" structs:"sort"`         //图片排序，例如2 其他图片可以是1 3
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" msg:"请完善菜单标题" structs:"menu_title"`          //菜单标题
	MenuTitleEn   string      `json:"menu_title_en" msg:"请完善菜单英文标题"  structs:"menu_title_en"` //英文菜单标题
	Slogan        string      `json:"slogan" structs:"slogan" `                               //广告
	Abstract      ctype.Array `json:"abstract"  structs:"abstract"`                           //简介
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time" `                 //简介的切换时间
	BannerTime    int         `json:"banner_time"  structs:"banner_time"`                     //菜单图片的切换时间
	Sort          int         `json:"sort"  msg:"请完善菜单排序" structs:"sort"`                     //菜单列表排序
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                            //具体图片的循序
}

// MenuCreateView 菜单创建
// @Tags 菜单管理
// @summary 菜单创建
// @Description 菜单创建
// @Param cr body MenuRequest true "菜单表的参数和关联图片表参数"
// @Router /menu [post]
// @Produce json
// @success 200 {object} res.Response
func (MenuAPi) MenuCreateView(c *gin.Context) {

	var cr MenuRequest
	//		绑定参数
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//绑定错误 返回msg tag 中错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	var menuModel = models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	//重复判断
	err = global.DB.Take(&menuModel, "menu_title = ?", menuModel.MenuTitle).Error
	if err == nil {
		res.FailWithMassage("标题重复", c)
		return
	}

	//菜单表入库
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		res.FailWithMassage("菜单入库失败", c)
		global.Log.Error(err)
		return
	}

	//判断是否有ImagesSort
	if len(cr.ImageSortList) == 0 {
		res.OkWith(c)
		return
	}

	var bannerImageModelList []models.MenuBannerModel
	//关联第三张表
	for _, sort := range cr.ImageSortList {
		//	这里应该有根据图片id判断是否有该图片
		bannerImageModelList = append(bannerImageModelList, models.MenuBannerModel{
			BannerID: sort.ImageID,
			MenuID:   menuModel.ID,
			Sort:     sort.Sort,
		})
	}
	//菜单表关联表入库
	//关联表中插入数据时，会判断外键关联的字段值是否存在
	err = global.DB.Create(&bannerImageModelList).Error
	if err != nil {
		res.FailWithMassage("菜单关联表入库失败", c)
		global.Log.Error(err)
		return
	}

	//	成功响应
	res.OkWith(c)

}
