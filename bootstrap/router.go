package bootstrap

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/middleware"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/router"
	"github.com/gin-gonic/gin"
)

// InitRouter 定义路由, init 版本只使用了两个中间件, TODO JWT
func InitRouter(r *gin.Engine) {
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 分组，路由前缀为"/douyin"
	dy := r.Group("/douyin")
	// 3.1 基础接口: 用户接口
	user := dy.Group("/user")
	// 用户注册、登录接口
	router.UserRouterGroup(user)
	// 3.1 基础接口：视频上传相关
	publish := dy.Group("/publish")
	// 视频投稿、发布列表
	router.PublishRouterGroup(publish)
}