package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
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
	list := make([]response.RelationUser, len(result))

	for i := 0; i < len(result); i++ {
		followId, _ := strconv.ParseInt(result[i], 10, 64)
		modelUser, _ := dao.UserInfoById(followId)
		list[i] = response.RelationUser{
			Id:            modelUser.Id,
			Name:          modelUser.Name,
			FollowCount:   dao.QueryFollowCount(c, result[i]),
			FollowerCount: dao.QueryFollowerCount(c, result[i]),
			IsFollow:      dao.IsFollow(c, userId, result[i]),
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
	list := make([]response.RelationUser, len(result))

	for i := 0; i < len(result); i++ {
		followId, _ := strconv.ParseInt(result[i], 10, 64)
		modelUser, _ := dao.UserInfoById(followId)
		list[i] = response.RelationUser{
			Id:            modelUser.Id,
			Name:          modelUser.Name,
			FollowCount:   dao.QueryFollowCount(c, result[i]),
			FollowerCount: dao.QueryFollowerCount(c, result[i]),
			IsFollow:      dao.IsFollow(c, userId, result[i]),
		}
	}

	c.JSON(http.StatusOK, response.RelationFollowListResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "获取用户粉丝列表成功"},
		UserList: list,
	})
}
