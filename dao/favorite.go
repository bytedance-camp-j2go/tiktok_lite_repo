package dao

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/gin-gonic/gin"
)

//提供方法：根据视频id查询出视频点赞数
func GetFavoriteCountByVideoId(c *gin.Context, videoId string) int64 {
	result, err := global.RedisDB.SCard(c, "favorite_count_set::"+videoId).Result()
	if err != nil {
		return 0
	}
	return result
}

//根据视频id查询是否已经点赞过
func IsFavorite(c *gin.Context, videoId string, userId string) bool {
	result, err := global.RedisDB.SIsMember(c, "favorite_count_set::"+videoId, userId).Result()
	if err != nil {
		return false
	}
	return result
}
