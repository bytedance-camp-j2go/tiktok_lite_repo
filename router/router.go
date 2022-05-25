package router

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/controller"
	"github.com/gin-gonic/gin"
)

func Feed(r *gin.RouterGroup) {
	r.GET("/feed/", controller.Feed)
}
