package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io/fs"
	"os"
	"path"
	"strings"
)

// imagesSuffixList 图片上传文件的后缀白名单
var imagesSuffixList = []string{
	"jpg",
	"jpeg",
	"png",
	"gif",
	"bmp",
	"tiff",
	"tif",
	"webp",
	"heif",
	"heic",
	"svg",
	"raw",
	"cr2",
	"nef",
	"arw",
	"psd",
	"ai",
	"eps",
	"ico",
	"pdf",
	"tga",
	"jp2",
	"j2k",
	"dds",
	"xcf",
}

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
	//判断下载路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {

		//不存在报错，就创建
		err = os.MkdirAll(basePath, fs.ModePerm) //Mkdir 方法不能直接创建多级目录
		if err != nil {
			res.FailWithMassage(err.Error(), c)
			return
		}
	}

	var fileResList []FileUploadResponse
	//	将每个文件单独判断是否可以上传 并将信息保存在fileUploadResponse中，最后统一响应
	for _, file := range files {
		//图片文件白名单
		//判断是否为图片格式后缀的文件
		filenameSplitList := strings.Split(file.Filename, ".")
		suffix := filenameSplitList[len(filenameSplitList)-1]
		exit := utils.InList(imagesSuffixList, suffix)
		if !exit {
			fileResList = append(fileResList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "文件非法",
			})
			continue
		}

		//	若超过系统限定大小 则上传失败
		size := float64(file.Size) / (float64(1024 * 1024)) //将字节转换为MB
		if size >= float64(global.Config.Upload.Size) {
			fileResList = append(fileResList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片上传失败，图片大小为： %.2f MB,系统图片上传阈值为： %d MB ", size, global.Config.Upload.Size),
			})
			continue
		}

		//上传图片
		err = c.SaveUploadedFile(file, path.Join(basePath, file.Filename))
		if err != nil {
			global.Log.Error(err)
			fileResList = append(fileResList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		//	上传成功
		fileResList = append(fileResList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}

	//响应成功或者失败的信息
	res.OkWithData(fileResList, c)
}
