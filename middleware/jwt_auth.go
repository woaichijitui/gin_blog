package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils"
)

// JwtAuth 普通用户授权中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否是管理者
		token := c.GetHeader("token")
		//有无token
		if token == "" {
			res.FailWithMassage("未携带token", c)
			c.Abort()
			return
		}
		//解析token
		claims, err := utils.ParseTokenRs256(token)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMassage("token解析失败", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)

	}
}

// JwtAdmin 管理用户授权中间件
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否是管理者
		token := c.GetHeader("token")
		//有无token
		if token == "" {
			res.FailWithMassage("未携带token", c)
			c.Abort()
			return
		}
		//解析token
		claims, err := utils.ParseTokenRs256(token)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMassage("token解析失败", c)
			c.Abort()
			return
		}

		//注销的用户
		logout := service.Service.RedisService.CheckLogout(token)
		if logout {
			res.FailWithMassage("用户已注销", c)
			c.Abort()
			return
		}

		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//	若不是管理员
			res.FailWithMassage("非管理员用户", c)
			return
		}

		c.Set("claims", claims)
	}

}
