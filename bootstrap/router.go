package bootstrap

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 定义路由, init 版本只使用了两个中间件, TODO JWT
func InitRouter(r *gin.Engine) {
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	dy := r.Group("/douyin")

	// TODO SOME IMPL

	dy.POST("/favorite/action", func(c *gin.Context) {
		userId, _ := c.GetPostForm("user_id")
		token, _ := c.GetPostForm("token")
		videoId, _ := c.GetPostForm("video_id")
		actionType, _ := c.GetPostForm("action_type")
		fmt.Println(userId, token, videoId, actionType)
		c.JSON(200, "点赞成功")
	})
}
