package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/controller"
)

func Feed(r *gin.RouterGroup) {
	// r.GET("/feed/", controller.Feed)
	r.GET("/feed", controller.Feed)
}
