package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"net/url"
	"os"
	"strings"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片的hash值，用于判断重复的图片
	Name      string          `json:"name" gorm:"size:256"'`       //图片名称
	ImageType ctype.ImageType `json:"image_type" gorm:"default:1"` //图片的类型，本地还是阿里云
}

// BeforeDelete 钩子函数将在删除记录前被调用
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {

	if b.ImageType == ctype.Local { //是本地文件
		//	本地图片，删除，还要删除本地的存储。本地文件和mysql中的banner表是一一对应的，无重复图片和本地图片
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	} else if b.ImageType == ctype.Aliyun {
		//判断阿里云buckey 是否为空，
		//为空则重新创建
		if global.Bucket == nil {
			//循环引用了
			/*bucket, err := service_com.CreateAliyunClient()
			//创建失败
			if err != nil {
				err = errors.New(fmt.Sprintf("连接阿里云失败:%v", err))
				global.Log.Error(err)
				return err
			}
			//若创建成功 赋予全局global中
			global.Bucket = bucket*/
			err = errors.New(fmt.Sprintf("连接阿里云失败:%v", err))
			global.Log.Error(err)
			return err

		}

		//获取得到阿里云文件空间文件的完整路径（不包含buckeyname
		// 解析URL
		parsedURL, err := url.Parse(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
		path := parsedURL.Path
		// 获取路径并去除开头的斜杠
		path = strings.TrimPrefix(parsedURL.Path, "/")
		fmt.Println(path)

		//	删除阿里云图片
		err = global.Bucket.DeleteObject(path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}

	//	删除成功
	global.Log.Info(fmt.Sprintf("删除照片,bannner id = %d", b.ID))

	return nil
}
