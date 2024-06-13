package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func Makemigrations() {
	var err error
	// 创建第三张连接表，many2many
	//	参数1：参数2：带外键的字段	参数3：连接表
	global.Mysql.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.Mysql.SetupJoinTable(&models.MenuModel{}, "user2_collects", &models.MenuBannerModel{})

	//	生成四张表的表结构
	err = global.Mysql.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.MenuModel{},
			&models.LoginDataModel{},
			&models.FadeBackModel{},
			&models.CommentModel{},
			&models.BannerModel{},
			&models.ArticleModel{},
			&models.AdvertModel{},
			//若是自定义第三章表 则需要手动迁移
			&models.MenuBannerModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Log.Infoln("[success] 生成数据库表结构成功！")
}
