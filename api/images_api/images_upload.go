package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/service_com"
	"gvb_server/utils"
	"gvb_server/utils/common"
	"io/ioutil"
	"path"
)

// FileUploadResponse 文件是否上传成功的响应模型
type FileUploadResponse struct {
	FileName  string `json:"file_name" `
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImagesUploadView 多个文件下载接口
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	//	接收多个文件（图片）
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMassage(err.Error(), c)
		return
	}
	files, ok := form.File["images"]
	if !ok {
		res.FailWithMassage("images文件不存在", c)
		return
	}

	var fileResList []res.FileUploadResponse
	//	将每个文件单独判断是否可以上传 并将信息保存在fileUploadResponse中，最后统一响应
	for _, file := range files {

		var fileName = file.Filename
		//图片文件白名单
		//判断是否为图片格式后缀的文件
		_, err := common.CheckFileSuffixIsRight(fileName)
		if err != nil {
			fileResList = append(fileResList, res.FileUploadResponse{
				FileName:  fileName,
				Url:       "",
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		//打开图片
		fileObj, err := file.Open()
		if err != nil {
			//记录发生错误,但不返回数据,也不跳过,因为前面可以确认该图片是正确的
			global.Log.Error(err)
			fileResList = append(fileResList, res.FileUploadResponse{
				FileName:  fileName,
				Url:       "",
				IsSuccess: false,
				Msg:       "文件打开失败",
			})
			continue
		}

		//计算图片的hash值
		data, err := ioutil.ReadAll(fileObj)
		if err != nil {
			//记录发生错误,但不返回数据,也不跳过,因为前面可以确认该图片是正确的
			global.Log.Error(err)
			fileResList = append(fileResList, res.FileUploadResponse{
				FileName:  fileName,
				Url:       "",
				IsSuccess: false,
				Msg:       "文件读取失败",
			})
			continue
		}
		hash := utils.Md5(data)

		var banner models.BannerModel
		//	查询数据是否有重复的图片
		row := global.DB.Take(&banner, "hash = ?", hash).RowsAffected
		if row != 0 {
			//	找到了
			fileResList = append(fileResList, res.FileUploadResponse{
				FileName:  fileName,
				Url:       banner.Path,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		//是否开启阿里云
		if global.Config.Aliyun.Enable {
			//	开启
			//	若超过系统限定大小 则上传失败
			err := common.CheckFileSizeOutOfLimit(file.Size, global.Config.Aliyun.Size)
			if err != nil {
				fileResList = append(fileResList, res.FileUploadResponse{
					FileName:  fileName,
					Url:       "",
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}

			//	上传阿里云
			//	fileheader
			fileHeader, err := file.Open()
			url, err := service_com.UploadFileAliyun(fileHeader, fileName)
			if err != nil {
				fileResList = append(fileResList, res.FileUploadResponse{
					FileName:  fileName,
					Url:       "",
					IsSuccess: false,
					Msg:       err.Error(),
				})
			}
			//	上传阿里云成功
			fileResList = append(fileResList, res.FileUploadResponse{
				FileName:  fileName,
				Url:       url,
				IsSuccess: true,
				Msg:       "上传阿里云成功",
			})

			//	存入数据库 数据库path 既阿里云url
			global.DB.Create(&models.BannerModel{
				Path:      url,
				Hash:      hash,
				Name:      fileName,
				ImageType: ctype.Aliyun,
			})
		} else {

			//上传本地

			basePath := global.Config.Upload.Path
			path := path.Join(basePath, fileName)
			err = common.UploadFileInLocal(data, path)
			if err != nil {
				global.Log.Error(err)
				fileResList = append(fileResList, res.FileUploadResponse{
					FileName:  fileName,
					Url:       path, //？
					IsSuccess: false,
					Msg:       err.Error(),
				})
				continue
			}
			// 上传本地成功
			fileResList = append(fileResList, res.FileUploadResponse{
				FileName:  fileName,
				Url:       path,
				IsSuccess: true,
				Msg:       "上传本地成功",
			})
			//	存入数据库
			global.DB.Create(&models.BannerModel{
				Path:      path,
				Hash:      hash,
				Name:      fileName,
				ImageType: ctype.Local,
			})
		}

	}

	//响应成功或者失败的信息
	res.OkWithData(fileResList, c)

}
