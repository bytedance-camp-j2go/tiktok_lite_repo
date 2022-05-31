package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/controller"
)

func RelationRouterGroup(relationGroup *gin.RouterGroup) {

	relationGroup.POST("/action/", controller.RelationAction)

	relationGroup.GET("/follow/list", controller.RelationFollowList)

	relationGroup.GET("/follower/list", controller.RelationFollowerList)
}
