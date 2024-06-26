package flag

import sys_flag "flag"

type Option struct {
	DB   bool
	User string //-u user or  -u admin
}

// Parse 解析命令行参数
//
//	是否初始化数据库
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库") //命令行参数若没加任何参数 则为true
	user := sys_flag.String("u", "", "创建权限用户")
	sys_flag.Parse()

	return Option{
		DB:   *db,
		User: *user,
	}

}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		//	若初始化数据库则停止
		return true
	}
	if option.User != "" {
		//	若初始化数据库则停止
		return true
	}

	return false

}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) bool {
	//调用迁移函数
	if option.DB {
		Makemigrations()
	}
	//调用创建用户函数
	if option.User == "user" || option.User == "admin" {
		CreateUser(option.User)
	}
	return true
}
