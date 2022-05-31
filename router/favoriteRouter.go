package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/controller"
)

func FavoriteRouterGroup(favoriteGroup *gin.RouterGroup) {

	favoriteGroup.POST("/action/", controller.FavoriteAction)

	favoriteGroup.GET("/list/", controller.FavoriteList)
}
