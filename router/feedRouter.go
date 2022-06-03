package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/controller"
)

func FeedRouterGroup(r *gin.RouterGroup) {
	// r.GET("/feed/", controller.FeedRouterGroup)
	r.GET("/feed", controller.Feed)
}
