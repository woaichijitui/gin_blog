package main

import (
	"fmt"
	"gvb_server/core"
	_ "gvb_server/docs" //需要导入文档包 不然报错
	//_ 表示执行init函数时调用改包
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/service/service_com"
)

// 添加注释以描述 server 信息 ，用于创建swagger文档
// @title           gin_blog API
// @version         1.0
// @description     gin_vue_blog 程序api
// @host      127.0.0.1:8080
// @BasePath  /api
func main() {
	//	读取配置文件
	core.InitConfig()

	//	初始化log
	global.Log = core.InitLogger()
	global.Log.Printf("初始化log成功，日志等级为：%s", global.Config.Logger.Lever)

	//	初始化mysql
	global.DB = core.InitGorm()
	global.Log.Infof(fmt.Sprintf("[%s] mysql连接成功！", global.Config.Mysql.DNS()))

	//
	global.Redis = core.ConnectRedisDB()
	if global.Redis == nil {
		global.Log.Error("redis连接失败")
		return
	}
	global.Log.Infof(fmt.Sprintf("[%s] redis连接成功！", global.Config.Redis.Addr()))

	//	命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		if flag.SwitchOption(option) {
			return
		}
	}

	//	启动gin
	addr := global.Config.System.Addr()
	router := routers.InitRouter()

	//连接aliyun
	service_com.AliyunInit()

	//	终端显示运行地址
	global.Log.Infof("gvb 运行在%s", addr)
	router.Run()
}
