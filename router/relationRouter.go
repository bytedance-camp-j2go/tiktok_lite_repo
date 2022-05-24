package router

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/controller"
	"github.com/gin-gonic/gin"
)

func RelationRouterGroup(relationGroup *gin.RouterGroup) {

	relationGroup.POST("/action/", controller.RelationAction)

	relationGroup.GET("/follow/list", controller.RelationFollowList)

	relationGroup.GET("/follower/list", controller.RelationFollowerList)
}
