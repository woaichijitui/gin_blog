package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		return
	}
	//	若没有该错误
	Result(ERROR, map[string]interface{}{}, "未知错误", c)
}
