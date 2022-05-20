/**
* @Author:drl
* @Date: 2022/5/18 23:49
 */
package router

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/controller"
	"github.com/gin-gonic/gin"
)

//用户路径下的所有请求  "/user","/user/register","/login"
func UserRouterGroup(userGroup *gin.RouterGroup) {
	//获取用户信息
	userGroup.GET("/", controller.UserController{}.User)
	//用户注册
	userGroup.POST("/register", controller.UserController{}.Register)
	//用户登录
	userGroup.POST("/login", controller.UserController{}.Login)

}
