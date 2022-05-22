package router

import "github.com/gin-gonic/gin"

// PublishRouterGroup publish下的所有请求：/publish/list/ /publish/action
func PublishRouterGroup(publishGroup *gin.RouterGroup) {
	//视频投稿
	publishGroup.POST("/action/")
	//发布列表
	publishGroup.GET("/list/")
}
