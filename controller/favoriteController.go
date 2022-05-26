package controller

import (
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
	userId := c.Query("user_id")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	switch actionType {
	case "1":
		zset := &redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: videoId,
		}
		// 维护一个排序集合，key为favorite_set::userId，value 为videoId，按照时间顺序排序
		err := global.RedisDB.ZAdd(c, "favorite_set::"+userId, zset).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis添加失败"},
			})
			return
		}

		// 维护一个set，key为favorite_count_set::videoId，value为用户id
		// 用来保存这个视频下面哪些用户点赞
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
		// 从zset中删除取消点赞的视频
		err := global.RedisDB.ZRem(c, "favorite_set::"+userId, videoId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FavoriteActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis zset 删除出错"},
			})
			return
		}

		// 取消点赞，从favorite_count_set::videoId删除该用户
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
	userId := c.Query("user_id")

	res, err := global.RedisDB.ZRange(c, "favorite_set::"+userId, 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FavoriteListResponse{
			Response:  response.Response{StatusCode: -1, StatusMsg: "redis查询出错"},
			VideoList: nil,
		})
		return
	}

	list := make([]model.Video, len(res))

	for i := 0; i < len(res); i++ {
		// todo 这里需要根据videoId获取视频对象
		videoId, _ := strconv.ParseInt(res[i], 10, 64)
		list[i] = model.Video{
			VideoId:       videoId,
			UserId:        videoId,
			PlayUrl:       "www.baidu.com",
			CoverUrl:      "www.taobao.com",
			FavoriteCount: 666,
			CommentCount:  222,
			Title:         res[i],
		}
	}

	c.JSON(http.StatusOK, response.FavoriteListResponse{
		Response:  response.Response{StatusCode: 0, StatusMsg: "获取列表成功"},
		VideoList: list,
	})

}

// 提供方法：根据视频id查询出视频点赞数
func GetFavoriteCountByVideoId(c *gin.Context, videoId string) int64 {
	result, err := global.RedisDB.SCard(c, "favorite_count_set::"+videoId).Result()
	if err != nil {
		return 0
	}
	return result
}

// 根据视频id查询是否已经点赞过
func IsFavorite(c *gin.Context, videoId string, userId string) bool {
	result, err := global.RedisDB.SIsMember(c, "favorite_count_set::"+videoId, userId).Result()
	if err != nil {
		return false
	}
	return result
}
