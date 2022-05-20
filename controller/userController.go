package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct{}

// UserInfo 获取用户信息
func UserInfo(context *gin.Context) {
	//分别获取发布视频者id和用户token
	//注意：这块使用context.query()获取的是字符串类型，需要将字符串转换为int64
	var publishId int64
	publishId, _ = strconv.ParseInt(context.Query("user_id"), 10, 64)
	//这块就不用鉴权了，能够进入到这块说明中间件那块已经鉴权过了，只需要获取用户信息
	//var user model.User
	user, _ := context.Get(global.UserName)
	//查询视频发布者信息
	publisher, _ := dao.UserInfoById(publishId)
	//查询用户是否已经关注这个视频发布者
	exists, _ := dao.UserFollower(user.(model.User).Id, publishId)
	//构建响应对象
	userResp := response.UserResponse{
		Response: response.Response{StatusCode: 200, StatusMsg: "成功"},
		User: response.User{
			Id:            publishId,
			Name:          publisher.Name,
			FollowCount:   publisher.FollowCount,
			FollowerCount: publisher.FollowerCount,
			IsFollow:      exists,
		},
	}
	context.JSON(http.StatusOK, userResp)
}

// UserRegister 用户注册
func UserRegister(context *gin.Context) {
	//获取账号密码
	//username := context.PostForm("username")
	//password := context.PostForm("password")
	username := context.Query("username")
	password := context.Query("password")
	//插入数据
	userId, err := dao.UserRegister(username, password)
	if err != nil {
		context.JSON(http.StatusBadRequest, "账号重复，请重新设置")
		return
	}
	//生成token
	user := model.User{Id: userId, UserName: username, PassWord: password}
	token, _ := util.GetToken(user)
	context.JSON(http.StatusOK, response.UserRegisterResponse{
		Response: response.Response{StatusCode: 200},
		UserId:   userId,
		Token:    token,
	})
}

// UserLogin 用户登录
func UserLogin(context *gin.Context) {
	//获取账号密码
	//username := context.PostForm("username")
	//password := context.PostForm("password")
	username := context.Query("username")
	password := context.Query("password")
	//查询用户是否存在
	user, err := dao.UserLogin(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, "账号或密码错误")
		return
	}
	//strings.Compare(s1,s2), 0代表相等，1代表s1>s2,-1代表s1<s2
	if strings.Compare(password, user.PassWord) == 0 {
		//获取token
		token, _ := util.GetToken(user)
		context.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 200, StatusMsg: "登录成功"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		context.JSON(http.StatusBadRequest, "账号或密码错误")
		return
	}
}
