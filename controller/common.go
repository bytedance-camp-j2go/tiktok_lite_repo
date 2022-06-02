package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tiktok-lite/dao"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/response"
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

// videoFeed 获取 video 信息并封装用户信息，封装
func videoFeed(videoIds []int64, uid int64) []response.Video {
	videos, err := dao.VideoQueryList(videoIds)
	if err != nil {
		zap.L().Error("获取视频信息失败!!", zap.Error(err))
		return nil
	}

	res := make([]response.Video, 0, len(videos))

	// 不在 for 循环内声明变量
	var author response.User
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
