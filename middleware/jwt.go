package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tiktok-lite/global"
	"tiktok-lite/util"
)

var (
	path     = []string{"feed", "login", "register"}
	whiteSet map[string]struct{}
)

func init() {
	whiteSet = make(map[string]struct{}, len(path))
	for _, key := range path {
		whiteSet[key] = struct{}{}
	}
}

// JWTAuth 鉴权中间件
// 注意：这个中间件不能直接初始化在main方法中，用户未登录可以刷视频，即可以获取到视频流，所以只需要在需要鉴权的路由中进行拦截
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的url，对其进行判断是否需要鉴权
		url := c.FullPath()
		strs := strings.Split(url, "/")
		inWhiteSet := false

		for _, s := range strs {
			if _, inWhiteSet = whiteSet[s]; inWhiteSet {
				// 能够进来的说明是不需要鉴权的接口，直接放行
				// c.Next()
				// return
				break
			}
		}

		// 获取请求头中 token，实际是一个完整被签名过的 token
		// tokenStr := c.GetHeader("Authorization")
		tokenStr := getToken(c)
		// fmt.Println(tokenStr)
		if tokenStr == "" && !inWhiteSet {
			c.JSON(http.StatusForbidden, "Permission denied!")
			c.Abort()
			return
		}

		// 解析token，并且获取用户信息
		claim, err := util.ParseToken(tokenStr)
		if err != nil {
			if inWhiteSet {
				c.Next()
				return
			}

			c.JSON(http.StatusForbidden, "Invalid token! You don't have permission!")
			c.Abort()
			return
		}

		// 将claim中的user信息存在context中
		c.Set(global.CtxUserKey, claim.User)

		// 这里执行路由 HandlerFunc
		c.Next()
	}
}

const JWTAuthKey = "token"

func getToken(ctx *gin.Context) (token string) {
	if token = ctx.Query(JWTAuthKey); token != "" {
		return token
	}
	if token = ctx.PostForm(JWTAuthKey); token != "" {
		return token
	}

	return ""
}
