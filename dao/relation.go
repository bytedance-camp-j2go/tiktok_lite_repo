package dao

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/gin-gonic/gin"
)

//方法：判断对方是否是我的关注
func IsFollow(c *gin.Context, userId string, toUserId string) bool {
	result, err := global.RedisDB.SIsMember(c, "follow_list::"+userId, toUserId).Result()
	if err != nil {
		return false
	}
	return result
}

//方法：查询我的关注数
func QueryFollowCount(c *gin.Context, userId string) int64 {
	result, err := global.RedisDB.SCard(c, "follow_list::"+userId).Result()
	if err != nil {
		return 0
	}
	return result
}

//方法：判断对方是否关注了我
func IsFollower(c *gin.Context, userId string, toUserId string) bool {
	result, err := global.RedisDB.SIsMember(c, "follower_list::"+userId, toUserId).Result()
	if err != nil {
		return false
	}
	return result
}

//方法：查询我的粉丝数
func QueryFollowerCount(c *gin.Context, userId string) int64 {
	result, err := global.RedisDB.SCard(c, "follower_list::"+userId).Result()
	if err != nil {
		return 0
	}
	return result
}
