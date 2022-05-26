package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
关注操作
*/
func RelationAction(c *gin.Context) {
	userId := c.Query("user_id")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	switch actionType {
	//关注操作
	case "1":
		err := global.RedisDB.SAdd(c, "follow_list::"+userId, toUserId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.RelationActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis sadd出错"},
			})
			return
		}
		err = global.RedisDB.SAdd(c, "follower_list::"+toUserId, userId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.RelationActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis sadd出错"},
			})
			return
		}
		c.JSON(http.StatusOK, response.RelationActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "关注成功"},
		})
	//取消关注
	case "2":
		err := global.RedisDB.SRem(c, "follow_list::"+userId, toUserId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.RelationActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis sadd出错"},
			})
			return
		}
		err = global.RedisDB.SRem(c, "follower_list::"+toUserId, userId).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.RelationActionResponse{
				Response: response.Response{StatusCode: -1, StatusMsg: "redis sadd出错"},
			})
			return
		}
		c.JSON(http.StatusOK, response.RelationActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "取消关注成功"},
		})
	//请求错误
	default:
		c.JSON(http.StatusBadRequest, response.RelationActionResponse{
			Response: response.Response{StatusCode: -1, StatusMsg: "行为异常"},
		})
		return
	}
}

/**
用户关注列表
*/
func RelationFollowList(c *gin.Context) {
	userId := c.Query("user_id")
	result, err := global.RedisDB.SMembers(c, "follow_list::"+userId).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.RelationActionResponse{
			Response: response.Response{StatusCode: -1, StatusMsg: "redis smember出错"},
		})
		return
	}
	list := make([]model.User, len(result))

	for i := 0; i < len(result); i++ {
		followId, _ := strconv.ParseInt(result[i], 10, 64)
		//todo 根据用户id查找用户方法
		list[i] = model.User{
			Id:            followId,
			UserName:      "liyu" + result[i],
			PassWord:      "123" + result[i],
			Name:          "liyu" + result[i],
			FollowCount:   100,
			FollowerCount: 100,
			IsFollow:      true,
		}
	}

	c.JSON(http.StatusOK, response.RelationFollowListResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "获取用户关注列表成功"},
		UserList: list,
	})

}

/**
用户粉丝列表
*/
func RelationFollowerList(c *gin.Context) {
	userId := c.Query("user_id")
	result, err := global.RedisDB.SMembers(c, "follower_list::"+userId).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.RelationActionResponse{
			Response: response.Response{StatusCode: -1, StatusMsg: "redis smember出错"},
		})
		return
	}
	list := make([]model.User, len(result))

	for i := 0; i < len(result); i++ {
		followerId, _ := strconv.ParseInt(result[i], 10, 64)
		// todo 根据用户id查找用户方法，这里暂时写死
		list[i] = model.User{
			Id:            followerId,
			UserName:      "liyu" + result[i],
			PassWord:      "123" + result[i],
			Name:          "liyu" + result[i],
			FollowCount:   100,
			FollowerCount: 100,
			IsFollow:      true,
		}
	}

	c.JSON(http.StatusOK, response.RelationFollowListResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "获取用户粉丝列表成功"},
		UserList: list,
	})
}

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
