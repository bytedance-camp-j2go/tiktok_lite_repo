package dao

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
)

//提供方法：根据视频id查询出视频点赞数
func GetFavoriteCountByVideoId(videoId string) int64 {
	result, err := global.RedisDB.SCard(global.RedisDB.Context(), "favorite_count_set::"+videoId).Result()
	if err != nil {
		return 0
	}
	return result
}

//根据视频id查询是否已经点赞过
func IsFavorite(videoId string, userId string) bool {
	result, err := global.RedisDB.SIsMember(global.RedisDB.Context(), "favorite_count_set::"+videoId, userId).Result()
	if err != nil {
		return false
	}
	return result
}
