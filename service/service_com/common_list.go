package service_com

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Logger bool
}

func ComList[T any](model T, option Option) (modelList []T, count int64, err error) {

	DB := global.DB
	//是否开启mysql日志
	if option.Logger {

		DB = DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})
	}

	query := DB.Where(&model)

	// 统计数量
	count = query.Select("id").Find(&modelList).RowsAffected

	//受前面影响 需要手动重新赋值
	query = DB.Where(&model)
	// 排序
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间从后往前排
	}
	// offset
	offset := (option.Page - 1) * option.Limit
	//	若小于0，则说明输出页数是错误的（小于等于0） 或者就是没有输入该数据
	if offset < 0 {
		offset = 0
	}

	//	分页查找数据
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&modelList).Error
	return modelList, count, err
}
