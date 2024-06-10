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

6.响应模型和预定义error

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

