package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MenuImage struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

// MenuResponse
type MenuResponse struct {
	models.MenuModel
	MenuImages []MenuImage
}

// MenuListView 菜单列表
// @Tags 菜单管理
// @summary 菜单列表
// @Description 菜单列表
// @Router /menu [get]
// @Produce json
// @success 200 {object} res.Response{data=[]MenuResponse{}}
func (MenuAPi) MenuListView(c *gin.Context) {
	//联合查询
	//	无需分页
	var menuList []models.MenuModel
	var menuIDList []uint
	//	按照sort排序查询menu,并将menu id 扫描至list
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	//	按照sort排序查询连接表
	var MenuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&MenuBanners, "menu_id in ?", menuIDList)

	var menuResList []MenuResponse
	//	查询menu对应banner
	for _, menu := range menuList {
		var menuImages []MenuImage
		var menuImage MenuImage
		for _, menuBanner := range MenuBanners {
			if menu.ID == menuBanner.MenuID {
				err := global.DB.Find(&models.BannerModel{}, "id = ?", menuBanner.BannerID).Select("id", "path").Scan(&menuImage).Error
				if err != nil {
					res.FailWithMassage("关联图片未找到", c)
					return
				}
				//	找到了就入该条菜单的图片列表中
				menuImages = append(menuImages, menuImage)
			}
		}
		// 装入容器
		menuResList = append(menuResList, MenuResponse{
			menu,
			menuImages,
		})
	}

	//	响应
	res.OkWithData(menuResList, c)
}
