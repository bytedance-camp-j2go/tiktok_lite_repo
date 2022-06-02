package dao

import (
	"tiktok-lite/global"
)

// 方法：判断对方是否是我的关注
func IsFollow(userId string, toUserId string) bool {
	result, err := global.RedisDB.SIsMember(global.RedisDB.Context(), "follow_list::"+userId, toUserId).Result()
	if err != nil {
		return false
	}
	return result
}

// 方法：查询我的关注数
func QueryFollowCount(userId string) int64 {
	result, err := global.RedisDB.SCard(global.RedisDB.Context(), "follow_list::"+userId).Result()
	if err != nil {
		return 0
	}
	return result
	// return util.ZSCard(global.VideoFavoriteUserSetKeySuffix + userId)
}

// 方法：判断对方是否关注了我
func IsFollower(userId string, toUserId string) bool {
	result, err := global.RedisDB.SIsMember(global.RedisDB.Context(), "follower_list::"+userId, toUserId).Result()
	if err != nil {
		return false
	}
	return result
}

// 方法：查询我的粉丝数
func QueryFollowerCount(userId string) int64 {
	result, err := global.RedisDB.SCard(global.RedisDB.Context(), "follower_list::"+userId).Result()
	if err != nil {
		return 0
	}
	return result
	// return util.ZSCard(global.VideoFavoriteUserSetKeySuffix + userId)
}
