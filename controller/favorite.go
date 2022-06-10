package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"net/http"
	"tiktok-lite/global"
	"tiktok-lite/response"
	"tiktok-lite/util"
)

// FavoriteAction 点赞行为
func FavoriteAction(c *gin.Context) {
	// 获取用户信息
	u := CtxUser(c)
	if u == DefUser {
		zap.L().Error("user info err!")
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 2, StatusMsg: "user login status error!!"})
		return
	}

	user := *u
	actionType := c.Query("action_type")
	//
	userId := user.Id
	videoId, err := util.String10Bit2Int64(c.Query("video_id"))
	if err != nil {
		zap.L().Debug("parse video id error !!", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.BaseInputError("输入错误的 video id 错误"))
	}

	UserFavoriteKey := fmt.Sprintf(global.UserFavoriteVideoSetKeySuffix, userId)
	VideoFavoriteKey := fmt.Sprintf(global.VideoFavoriteUserSetKeySuffix, videoId)

	switch actionType {
	case "1":
		go util.ZAdd2Redis(UserFavoriteKey, float64(util.TimeNowInt64()), videoId)
		go util.ZAdd2Redis(VideoFavoriteKey, float64(util.TimeNowInt64()), userId)
		c.JSON(http.StatusOK, response.BaseSuccess("点赞成功"))
	case "2":
		// 从 zset 中删除取消点赞的视频
		if util.ZRM2Redis(UserFavoriteKey, videoId) == 0 {
			c.JSON(http.StatusInternalServerError, response.BaseServerError("rm redis key error!! "+UserFavoriteKey))
			return
		}

		// 取消点赞，从favorite_count_set::videoId删除该用户
		if util.ZRM2Redis(VideoFavoriteKey, userId) == 0 {
			c.JSON(http.StatusInternalServerError, response.BaseServerError("rm redis key error!! "+VideoFavoriteKey))
			return
		}

		c.JSON(http.StatusOK, response.BaseSuccess("成功取消点赞"))

	default:
		c.JSON(http.StatusBadRequest, response.BaseInputError("非法点赞行为"))
	}

}

var maxRangeBy = &redis.ZRangeBy{
	Min: "-inf",
	Max: "+inf",
}

// FavoriteList 获取用户点赞列表
func FavoriteList(c *gin.Context) {
	user := *CtxUser(c)
	videoIdList, err := util.ZSetRangeByScoreInt(fmt.Sprintf(global.UserFavoriteVideoSetKeySuffix, user.Id), maxRangeBy)
	if len(videoIdList) == 0 && err != nil {
		zap.L().Debug("get favorite video list err!!", zap.Int("len", len(videoIdList)), zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	VideoFeed(videoIdList, user.Id)
	c.JSON(http.StatusOK, response.FavoriteListResponse{
		Response:  response.BaseSuccess("get publish list success"),
		VideoList: VideoFeed(videoIdList, user.Id),
	})
}
