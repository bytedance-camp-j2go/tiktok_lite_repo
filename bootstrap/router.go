package bootstrap

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/middleware"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/router"
	"github.com/gin-gonic/gin"
)

// Router 定义路由, init 版本只使用了两个中间件, TODO JWT
func Router(r *gin.Engine) {
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	//分组，路由前缀为"/douyin"
	dy := r.Group("/douyin")
	//3.1 基础接口: 用户接口
	user := dy.Group("/user")
	//用户注册、登录接口
	router.UserRouterGroup(user)
}
