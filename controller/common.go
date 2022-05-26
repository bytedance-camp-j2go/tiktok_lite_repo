package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/gin-gonic/gin"
)

var DefUser = &model.User{
	Id: -1,
}

// CtxUser 封装从 ctx 取 User 的方法，需要调用者判断指针是否等于 DefUser
func CtxUser(context *gin.Context) *model.User {
	// 从上下文获取用户信息
	ctxVal, _ := context.Get(global.CtxUserKey)

	user, ok := ctxVal.(model.User)
	if !ok {
		// zap.L().Error("user info err!")
		// context.JSON(http.StatusBadRequest, response.Response{StatusCode: 2, StatusMsg: ""})
		return DefUser
	}
	return &user
}
