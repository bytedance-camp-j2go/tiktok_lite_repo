package router

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/controller"
)

// UserRouterGroup 用户路径下的所有请求  "/user","/user/register","/login"
func UserRouterGroup(userGroup *gin.RouterGroup) {
	// 获取用户信息
	userGroup.GET("/", controller.UserInfo)
	// 用户注册
	userGroup.POST("/register/", controller.UserRegister)
	// 用户登录
	userGroup.POST("/login/", controller.UserLogin)

}
