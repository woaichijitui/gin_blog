## gin_vue_blog

### 一、项目搭建

#### 1.加载配置文件

#### 2.日志初始化

#### 3.数据库初始化

#### 4.路由初始化

enter

```go
func InitRouter() *gin.Engine {
	//设置gin的模式
	gin.SetMode(global.Config.System.Env)

	router := gin.Default()
	apiGroup := router.Group("/api")
	//系统设置路由
	SettingRouter(apiGroup)
	return router
}
```

路由分组

```go
func SettingRouter(router *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/setting", settingsApi.SettingInfoView)
}
```

![image-20240604215711058](C:\Users\hil\AppData\Roaming\Typora\typora-user-images\image-20240604215711058.png)

#### 5.api编写

![image-20240604215819855](C:\Users\hil\AppData\Roaming\Typora\typora-user-images\image-20240604215819855.png)

```
type SettingsApi struct {
}
```

```go
func (SettingsApi) SettingInfoView(c *gin.Context) {

	res.Ok(map[string]string{"name": "htt"}, "成功", c)

}
```

通过嵌入结构体，调用结构体方法（handle），调用api 。函数式编程。

```go
type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)

```

#### 6.错误状态码

![image-20240604220412868](C:\Users\hil\AppData\Roaming\Typora\typora-user-images\image-20240604220412868.png)

```go
const (
	SUCCESS = 0
	ERROR   = 7
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithMassage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMassage(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

// 根据code 查询出msg
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, b := ErrorMap[code]
	// 若有该错误，则取其内容
	if b {
		Result(int(code), map[string]interface{}{}, msg, c)
	}
	//	若没有该错误
	Result(ERROR, map[string]interface{}{}, "未知错误", c)
}
```

通过映射关系 ，一个code对应一个错误信息

```go
type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //系统错误
)

var ErrorMap = map[ErrorCode]string{
	SettingsError: "系统错误",
}

```

### 二、表结构搭建

![](C:\Users\hil\AppData\Roaming\Typora\typora-user-images\image-20240611131750808.png)

#### 表结构说明

```
advert_model.go 	//广告表
article_model.go	//文章表
banner_model.go		//图片表
comment_model.go	//评论表
enter.go	
fade_back_model.go	//用户反馈表
login_data_model.go	//登录信息表
menu_banner_model.go//菜单图片表
menu_model.go		//菜单表
message_model.go	//信息表
tag_model.go		//标签表
user_collect_model.go	//用户收藏文章表
user_model.go		//用户表
```

### 三、配置API

主要配置有：

site_info

qq

email

七牛

jwt

修改配置文件信息api

```go
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	//绑定json参数
	var siteInfo config.SiteInfo
	err := c.ShouldBindJSON(&siteInfo)
	if err != nil {
		global.Log.Errorln(err)
		//若是绑定失败 则返回失败信息
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//绑定成功
	//	改变与yaml文件绑定结构体
	global.Config.SiteInfo = siteInfo
	//	更改配置文件
	err = core.SetYaml()
	if err != nil {
		global.Log.Errorln(err)
		res.FailWithMassage(err.Error(), c)
	}
	res.OkWithMassage("更新成功", c)
	global.Log.Infoln("更改系统信息成功")

}

```

将多个配置浓缩为一个配置？



### 四、图片管理API

#### 1、图片上传view

```go
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	//	接收多个文件（图片）
	form, err := c.MultipartForm()
    files, ok := form.File["images"]
	//	将文件列表逐一判断并上传
    for _, file := range files {
        // ...
    }
    //	最后响应包含所有图片是否上传成功及其原因的res
    res.OkWithData(fileResList, c)
}
```

```go
type FileUploadResponse struct {
	FileName  string `json:"file_name" `
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}
```

```
//	FileUploadResponse 列表
var fileResList []FileUploadResponse
```

下载图片路径

```
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
```



#### 2、图片上传白名单

```
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
```

#### 3、图片存入数据库

`global.Mysql.Take(&models.BannerModel{}, "hash = ?", hash).RowsAffected`

```
//  存入数据库
global.Mysql.Create(&models.BannerModel{
    Path: path,
    Hash: hash,
    Name: file.Filename,
})
```

#### 4、获取图片list

```
// 图片列表查询接口
func (ImagesApi) ImagesListView(c *gin.Context) {
	var imagesList []models.BannerModel
	var page models.PageInfo
	//绑定参数
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	imagesList, count, _ := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: page,
		Logger:   true,
	})

	//	成功响应
	res.OkWithList(imagesList, count, c)

}
```

list的通用响应

```
// OkWithList 响应分页操作
func OkWithList[T any](list T, count int64, c *gin.Context) {
	OkWithData(gin.H{"list": list, "count": count}, c)
}
```

通用的分页查询

```
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

	// 统计数量
	count = DB.Select("id").Find(&modelList).RowsAffected
	// offset
	offset := (option.Page - 1) * option.Limit
	//	若小于0，则说明输出页数是错误的（小于等于0） 或者就是没有输入该数据
	if offset < 0 {
		offset = 0
	}

	//	分页查找数据
	err = DB.Limit(option.Limit).Offset(offset).Find(&modelList).Error
	return modelList, count, err
}
```



#### 5、图片删除操作

删除使用了Hook函数

```go
// BeforeDelete 钩子函数将在删除记录前被调用
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {

	if b.ImageType == ctype.Local { //是本地文件
		//	本地图片，删除，还要删除本地的存储。本地文件和mysql中的banner表是一一对应的，无重复图片和本地图片
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}

	return nil
}
```



