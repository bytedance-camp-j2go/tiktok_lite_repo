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
	userId := user.UserId
	videoId, err := util.String10Bit2Int64(c.Query("video_id"))
	if err != nil {
		zap.L().Debug("parse video id error !!", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.BaseInputError("输入错误的 video id 错误"))
	}

	// FavoriteSetKeySuffix := "favorite_set::"
	// FavoriteCntSetKeySuffix := "favorite_count_set::"
	// zap.L().Debug("userid >>", zap.Int64("uid", userId))
	UserFavoriteKey := fmt.Sprintf(global.UserFavoriteVideoSetKeySuffix, userId)
	VideoFavoriteKey := fmt.Sprintf(global.VideoFavoriteUserSetKeySuffix, videoId)

	switch actionType {
	case "1":
		// zset := &redis.Z{
		// 	Score:  float64(time.Now().Unix()),
		// 	Member: videoId,
		// }
		// // 维护一个排序集合，key为favorite_set::userId，value 为videoId，按照时间顺序排序
		// err := global.RedisDB.ZAdd(c, global.UserFavoriteVideoSetKeySuffix+userIdStr, zset).Err()

		// 等效

		go util.ZAdd2Redis(UserFavoriteKey, float64(util.TimeNowInt64()), videoId)

		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
		// 		Response: response.Response{StatusCode: -1, StatusMsg: "redis添加失败"},
		// 	})
		// 	return
		// }
		go util.ZAdd2Redis(VideoFavoriteKey, float64(util.TimeNowInt64()), userId)

		// // 维护一个set，key为favorite_count_set::videoId，value为用户id
		// // 用来保存这个视频下面哪些用户点赞
		// err = global.RedisDB.SAdd(c, FavoriteCntSetKeySuffix+videoId, userId).Err()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
		// 		Response: response.Response{StatusCode: -1, StatusMsg: "redis sadd出错"},
		// 	})
		// 	return
		// }

		// c.JSON(http.StatusOK, response.FavoriteActionResponse{
		// 	Response: response.Response{StatusCode: 0, StatusMsg: "点赞成功"},
		// })
		c.JSON(http.StatusOK, response.BaseSuccess("点赞成功"))
	case "2":
		// 从 zset 中删除取消点赞的视频
		if util.ZRM2Redis(UserFavoriteKey, videoId) == 0 {
			c.JSON(http.StatusInternalServerError, response.BaseServerError("rm redis key error!! "+UserFavoriteKey))
			return
		}

		// err := global.RedisDB.ZRem(c, global.UserFavoriteVideoSetKeySuffix+userIdStr, videoId).Err()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
		// 		Response: response.Response{StatusCode: -1, StatusMsg: "redis zset 删除出错"},
		// 	})
		// 	return
		// }

		// 取消点赞，从favorite_count_set::videoId删除该用户
		if util.ZRM2Redis(VideoFavoriteKey, userId) == 0 {
			c.JSON(http.StatusInternalServerError, response.BaseServerError("rm redis key error!! "+VideoFavoriteKey))
			return
		}
		// err = global.RedisDB.SRem(c, global.VideoFavoriteUserSetKeySuffix+videoId, userIdStr).Err()
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
		// 		Response: response.Response{StatusCode: -1, StatusMsg: "redis set 删除出错"},
		// 	})
		// 	return
		// }

		c.JSON(http.StatusOK, response.BaseSuccess("成功取消点赞"))
		// c.JSON(http.StatusOK, response.FavoriteActionResponse{
		// 	Response: response.Response{StatusCode: 0, StatusMsg: "成功取消点赞"},
		// })
	default:
		c.JSON(http.StatusBadRequest, response.BaseInputError("非法点赞行为"))
		// c.JSON(http.StatusBadRequest, response.FavoriteActionResponse{
		// 	Response: response.Response{StatusCode: -1, StatusMsg: "非法点赞行为"},
		// })
	}

}

var maxRangeBy = &redis.ZRangeBy{
	Min: "-inf",
	Max: "+inf",
}

// FavoriteList 获取用户点赞列表
func FavoriteList(c *gin.Context) {
	// // 获取用户信息
	// var a any
	// a, _ = c.Get(global.CtxUserKey)
	// user := a.(model.User)
	// userId := model.User(user).Id
	// userIdStr := strconv.FormatInt(userId, 10)
	//
	// res, err := global.RedisDB.ZRange(c, "favorite_set::"+userIdStr, 0, -1).Result()
	//
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, response.FavoriteListResponse{
	// 		Response:  response.Response{StatusCode: -1, StatusMsg: "redis查询出错"},
	// 		VideoList: nil,
	// 	})
	// 	return
	// }
	//
	// VideoIdInt64 := make([]int64, len(res))
	// for i := 0; i < len(res); i++ {
	// 	VideoIdInt64[i], _ = util.String10Bit2Int64(res[i])
	// }
	// // 获取视频对象数组
	// queryList, err := dao.VideoQueryList(VideoIdInt64)
	//
	// list := make([]response.FavoriteVideo, len(res))
	// // 封装视频数组
	// for i := 0; i < len(res); i++ {
	// 	userModel, _ := dao.UserInfoById(queryList[i].UserId)
	// 	videoUser := strconv.FormatInt(queryList[i].UserId, 10)
	// 	list[i] = response.FavoriteVideo{
	// 		VideoId: queryList[i].VideoId,
	// 		User: response.FavoriteUser{
	// 			Id:            userModel.Id,
	// 			Name:          userModel.Name,
	// 			FollowerCount: userModel.FollowerCount,
	// 			FollowCount:   userModel.FollowCount,
	// 			IsFollow:      dao.IsFollow(userIdStr, videoUser),
	// 		},
	// 		PlayUrl:       queryList[i].PlayUrl,
	// 		CoverUrl:      queryList[i].CoverUrl,
	// 		FavoriteCount: dao.GetFavoriteCountByVideoId(res[i]),
	// 		CommentCount:  queryList[i].CommentCount,
	// 		IsFavorite:    dao.IsFavorite(queryList[i].VideoId, userId),
	// 		Title:         queryList[i].Title,
	// 	}
	// }
	//
	// c.JSON(http.StatusOK, response.FavoriteListResponse{
	// 	Response:  response.Response{StatusCode: 0, StatusMsg: "获取列表成功"},
	// 	VideoList: list,
	// })

	user := *CtxUser(c)
	// 根据时间戳, 返回 list
	// feedProcess(c, start, user)
	videoIdList, err := util.ZSetRangeByScoreInt(fmt.Sprintf(global.UserFavoriteVideoSetKeySuffix, user.Id), maxRangeBy)
	if len(videoIdList) == 0 && err != nil {
		zap.L().Debug("get favorite video list err!!", zap.Int("len", len(videoIdList)), zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	videoFeed(videoIdList, user.Id)
	c.JSON(http.StatusOK, response.FavoriteListResponse{
		Response:  response.BaseSuccess("get publish list success"),
		VideoList: videoFeed(videoIdList, user.Id),
	})
}
