package service_com

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gvb_server/global"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// 上传文件到阿里云,并返回该图片的url 和 err
func UploadFileAliyun(file multipart.File, fileName string) (url string, err error) {

	//路径：配置路径+文件名
	Path := filepath.Join(global.Config.Aliyun.Prefix, fileName)
	// 替换反斜杠为正斜杠
	normalizedPath := strings.ReplaceAll(Path, "\\", "/")

	//判断阿里云buckey 是否为空，
	//为空则重新创建
	if global.Bucket == nil {
		bucket, err := CreateAliyunClient()
		//创建失败
		if err != nil {
			err = errors.New(fmt.Sprintf("连接阿里云失败:%v", err))
			return "", err
		}
		global.Bucket = bucket

	}

	// 上传文件到OSS
	//	上传相同名字的文件会覆盖之前上传的文件
	err = global.Bucket.PutObject(normalizedPath, file)
	if err != nil {
		fmt.Sprintf("Failed to upload file to OSS: %v", err)
		return "", err
	}

	//上传成功
	// 生成文件的URL
	//	外网访问URL：http://bucketexample.oss-cn-hangzhou.aliyuncs.com/example/example.jpg
	//	未开启https
	fmt.Println(global.Config.Aliyun.Endpoint)
	urlEndPoint := strings.Split(global.Config.Aliyun.Endpoint, "https://")[1]
	if urlEndPoint == "" {
		return "", errors.New("阿里云地域获取失败")
	}
	fileURL := fmt.Sprintf("http://%s.%s/%s", global.Config.Aliyun.BucketName, urlEndPoint, normalizedPath)
	return fileURL, nil
}

// 将buckey投入全局变量中
func AliyunInit() {
	if global.Config.Aliyun.Enable {
		bucket, err := CreateAliyunClient()
		if err != nil {
			//	阿里云创建失败
			global.Log.Error("阿里云未能连接成功： ", err)
			return
		}

		//创建成功
		global.Bucket = bucket

	}
}

// 创建aliyun客户端和容器空间
//
//	频繁创建是否影响性能？
func CreateAliyunClient() (bucket *oss.Bucket, err error) {
	// 创建OSS客户端
	client, err := oss.New(global.Config.Aliyun.Endpoint, global.Config.Aliyun.AccessKeyId, global.Config.Aliyun.AccessKeySecret)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to create OSS client: %v", err))
		return nil, err
	}

	// 获取存储空间
	bucket, err = client.Bucket(global.Config.Aliyun.BucketName)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to get OSS bucket: %v", err))
		return nil, err

	}
	return bucket, err

}
