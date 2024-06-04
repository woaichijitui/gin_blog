package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//	读取配置文件
	core.InitConfig()

	//	初始化log
	global.Log = core.InitLogger()
	global.Log.Printf("初始化log成功，日志等级为：%s", global.Config.Logger.Lever)

	//	初始化mysql
	global.Mysql = core.InitGorm()
	global.Log.Infof(fmt.Sprintf("[%s] mysql连接成功！", global.Config.Mysql.DNS()))

	//	启动gin
	addr := global.Config.System.Addr()
	router := routers.InitRouter()
	global.Log.Infof("gvb 运行在%s", addr)
	router.Run()
}
