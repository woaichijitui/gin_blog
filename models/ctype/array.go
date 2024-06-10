package ctype

import (
	"database/sql/driver"
	"strings"
)

/*
	数据库无直接存入go中的切片
	Scan 将数据（字符串）用"\n"切割成字符串切片返回给Array
	Value 将字符串切片用"\n"组合成一个字符串，并返回一个数据库可以识别类型
*/

type Array []string

func (a *Array) Scan(value interface{}) error {

	v, _ := value.([]byte)
	if string(v) == "" {
		*a = []string{}
		return nil
	}
	*a = strings.Split(string(v), "\n")
	return nil
}

func (a *Array) Value() (driver.Value, error) {
	//	将数值转化为值
	return strings.Join(*a, "\n"), nil
}
