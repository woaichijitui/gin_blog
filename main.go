package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//	读取配置文件
	core.InitConfig()

	//	初始化log
	global.Log = core.InitLogger()
	global.Log.Warnln("Warnln")
	global.Log.Errorln("Errorln")
	global.Log.Infoln("Infoln")
	global.Log.Debugln("Debugln")

	//	初始化mysql
	global.Mysql = core.InitGorm()
	fmt.Println(global.Mysql)

}
