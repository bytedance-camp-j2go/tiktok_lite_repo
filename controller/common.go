package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tiktok-lite/dao"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/response"
)

const (
	_ = iota
	ActionAppend
	ActionDel
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

func Videos2Response(videos []model.Video, uid int64) []response.Video {
	res := make([]response.Video, 0, len(videos))

	// 不在 for 循环内声明变量
	var (
		author response.User
		err    error
	)
	for _, video := range videos {
		author, err = response.NewUser(video.UserId, uid)
		if err != nil {
			zap.L().Error("get UserInfo", zap.Error(err))
			res = append(res, response.Video{Video: video})
			continue
		}
		res = append(res, response.Video{
			Video:      video,
			Author:     author,
			IsFavorite: dao.IsFavorite(video.VideoId, uid),
		})
	}
	return res
}

// VideoFeed 获取 video 信息并封装用户信息，封装
func VideoFeed(videoIds []int64, uid int64) []response.Video {
	videos, err := dao.VideoQueryList(videoIds)
	if err != nil {
		zap.L().Error("获取视频信息失败!!", zap.Error(err))
		return nil
	}

	return Videos2Response(videos, uid)
}

// 封装可复用 ctx 简单的处理方法

// CtxInputError 输入错误
func CtxInputError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, response.BaseInputError(msg))
}

func CtxServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, response.BaseServerError(msg))
}

func CtxBaseSuccess(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, response.BaseSuccess(msg))
}
