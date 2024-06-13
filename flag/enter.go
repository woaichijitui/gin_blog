package flag

import sys_flag "flag"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
//
//	是否初始化数据库
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")

	sys_flag.Parse()

	return Option{
		DB: *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		//	若初始化数据库则停止
		return true
	}

	return false

}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) bool {
	if option.DB {
		Makemigrations()
		return true
	}
	return false
}
