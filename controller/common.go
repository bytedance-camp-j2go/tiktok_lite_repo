package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok-lite/global"
	"tiktok-lite/model"
)

var DefUser = &model.User{
	Id: -1,
}

// CtxUser 封装从 ctx 取 User 的方法，需要调用者判断指针是否等于 DefUser
func CtxUser(context *gin.Context) *model.User {
	// 从上下文获取用户信息
	ctxVal, _ := context.Get(global.CtxUserKey)
	// context.Request.Host
	user, ok := ctxVal.(model.User)
	if !ok {
		// zap.L().Error("user info err!")
		// context.JSON(http.StatusBadRequest, response.Response{StatusCode: 2, StatusMsg: ""})
		return DefUser
	}
	return &user
}
