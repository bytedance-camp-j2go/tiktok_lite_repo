package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/controller"
)

func CommentRouterGroup(r *gin.RouterGroup) {
	r.POST("/action/", controller.CommentAction)
	r.GET("/list/", controller.CommentList)
}
