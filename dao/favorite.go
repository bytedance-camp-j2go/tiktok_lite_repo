package dao

import (
	"fmt"
	"tiktok-lite/global"
	"tiktok-lite/util"
)

/* Redis 操作，查询用户关注状态 */

// GetFavoriteCountByVideoId 提供方法：根据视频id查询出视频点赞数
func GetFavoriteCountByVideoId(videoId int64) int64 {
	return util.ZSetCnt(fmt.Sprintf(global.VideoFavoriteUserSetKeySuffix, videoId))
}

// IsFavorite 根据视频id查询是否已经点赞过
func IsFavorite(videoId, userId int64) bool {
	rank := util.ZSetRank(
		fmt.Sprintf(global.VideoFavoriteUserSetKeySuffix, videoId),
		util.Int64D2String(userId),
	)
	return rank > -1
}
