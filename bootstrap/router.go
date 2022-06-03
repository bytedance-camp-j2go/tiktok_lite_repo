package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/middleware"
	"tiktok-lite/router"
)

// InitRouter 定义路由, init 版本只使用了两个中间件
func InitRouter(r *gin.Engine) {
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 分组，路由前缀为"/douyin"
	dy := r.Group("/douyin")

	// 1. 基础接口: 用户接口组
	// 用户注册、登录
	user := dy.Group("/user")
	router.UserRouterGroup(user)

	// 2. 基础接口：视频发布接口组
	// 视频投稿、发布列表
	publish := dy.Group("/publish")
	router.PublishRouterGroup(publish)

	// 3. 视频流接口
	router.FeedRouterGroup(dy)

	// 4. 点赞接口组
	favorite := dy.Group("/favorite")
	router.FavoriteRouterGroup(favorite)

	// 5. 关注接口组
	relation := dy.Group("/relation")
	router.RelationRouterGroup(relation)

	// 6. 评论接口组
	comment := dy.Group("/comment")
	router.CommentRouterGroup(comment)
}
