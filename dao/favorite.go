package dao

import (
	"fmt"
	"tiktok-lite/global"
	"tiktok-lite/util"
)

/* Redis 操作，查询用户关注状态 */

// GetFavoriteCountByVideoId 提供方法：根据视频id查询出视频点赞数
func GetFavoriteCountByVideoId(videoId int64) int64 {
	// result, err := global.RedisDB.SCard(global.RedisDB.Context(), "favorite_count_set::"+videoId).Result()
	// if err != nil {
	// 	return 0
	// }
	// return result
	return util.ZSetCnt(fmt.Sprintf(global.VideoFavoriteUserSetKeySuffix, videoId))
}

// IsFavorite 根据视频id查询是否已经点赞过
func IsFavorite(videoId, userId int64) bool {
	result, err := global.RedisDB.SIsMember(
		global.RedisDB.Context(),
		fmt.Sprintf(global.VideoFavoriteUserSetKeySuffix, videoId),
		userId,
	).Result()
	if err != nil {
		return false
	}
	return result
}
