package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"
)

/*
点赞行为
*/
func FavoriteAction(c *gin.Context) {
	// 获取用户信息
	var a any
	a, _ = c.Get(global.CtxUserKey)
	user := a.(model.User)
	userId := model.User(user).Id
	userIdStr := strconv.FormatInt(userId, 10)

	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	switch actionType {
	case "1":
		zset := &redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: videoId,
		}
		//维护一个排序集合，key为favorite_set::userId，value 为videoId，按照时间顺序排序
		err := global.RedisDB.ZAdd(c, "favorite_set::"+userIdStr, zset).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis添加失败"},
			})
			return
		}

		//维护一个set，key为favorite_count_set::videoId，value为用户id
		//用来保存这个视频下面哪些用户点赞
		err = global.RedisDB.SAdd(c, "favorite_count_set::"+videoId, userId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis sadd出错"},
			})
			return
		}

		c.JSON(http.StatusOK, response.FavoriteActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "点赞成功"},
		})
	case "2":
		//从zset中删除取消点赞的视频
		err := global.RedisDB.ZRem(c, "favorite_set::"+userIdStr, videoId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis zset 删除出错"},
			})
			return
		}

		//取消点赞，从favorite_count_set::videoId删除该用户
		err = global.RedisDB.SRem(c, "favorite_count_set::"+videoId, userId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis set 删除出错"},
			})
			return
		}

		c.JSON(http.StatusOK, response.FavoriteActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "成功取消点赞"},
		})
	default:
		c.JSON(http.StatusBadRequest, response.FavoriteActionResponse{
			Response: response.Response{StatusCode: -1, StatusMsg: "非法点赞行为"},
		})
	}

}

/**
获取点赞列表
*/
func FavoriteList(c *gin.Context) {
	// 获取用户信息
	var a any
	a, _ = c.Get(global.CtxUserKey)
	user := a.(model.User)
	userId := model.User(user).Id
	userIdStr := strconv.FormatInt(userId, 10)

	res, err := global.RedisDB.ZRange(c, "favorite_set::"+userIdStr, 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FavoriteListResponse{
			Response:  response.Response{StatusCode: -1, StatusMsg: "redis查询出错"},
			VideoList: nil,
		})
		return
	}
	VideoIdInt64 := make([]int64, len(res))
	for i := 0; i < len(res); i++ {
		VideoIdInt64[i], _ = strconv.ParseInt(res[i], 10, 64)
	}
	//获取视频对象数组
	queryList, err := dao.VideoQueryList(VideoIdInt64)

	list := make([]response.FavoriteVideo, len(res))
	//封装视频数组
	for i := 0; i < len(res); i++ {
		userModel, _ := dao.UserInfoById(queryList[i].UserId)
		videoUser := strconv.FormatInt(queryList[i].UserId, 10)
		list[i] = response.FavoriteVideo{
			VideoId: queryList[i].VideoId,
			User: response.FavoriteUser{
				Id:            userModel.Id,
				Name:          userModel.Name,
				FollowerCount: userModel.FollowerCount,
				FollowCount:   userModel.FollowCount,
				IsFollow:      dao.IsFollow(userIdStr, videoUser),
			},
			PlayUrl:       queryList[i].PlayUrl,
			CoverUrl:      queryList[i].CoverUrl,
			FavoriteCount: dao.GetFavoriteCountByVideoId(res[i]),
			CommentCount:  queryList[i].CommentCount,
			IsFavorite:    dao.IsFavorite(queryList[i].VideoId, userId),
			Title:         queryList[i].Title,
		}
	}

	c.JSON(http.StatusOK, response.FavoriteListResponse{
		Response:  response.Response{StatusCode: 0, StatusMsg: "获取列表成功"},
		VideoList: list,
	})

}
