/**
* @Author:drl
* @Date: 2022/5/19 14:42
 */
package middleware

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuth 鉴权中间件
//注意：这个中间件不能直接初始化在main方法中，用户未登录可以刷视频，即可以获取到视频流，所以只需要在需要鉴权的路由中进行拦截
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中 token，实际是一个完整被签名过的 token
		tokenStr := c.GetHeader("Authorization")
		// fmt.Println(tokenStr)
		if tokenStr == "" {
			c.JSON(http.StatusForbidden, "无权访问，请求未带token")
			c.Abort()
			return
		}

		//解析token，并且获取用户信息
		claim, err := util.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusForbidden, "Invalid token! You don't have permission!")
			c.Abort()
			return
		}

		// 将claim中的user信息存在context中
		c.Set("userName", claim.User)

		// 这里执行路由 HandlerFunc
		c.Next()
	}
}
