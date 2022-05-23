package router

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/controller"
	"github.com/gin-gonic/gin"
)

func FavoriteRouterGroup(favoriteGroup *gin.RouterGroup) {

	favoriteGroup.POST("/action/", controller.FavoriteAction)

	favoriteGroup.GET("/list/", controller.FavoriteList)
}
