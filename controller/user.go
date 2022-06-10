package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tiktok-lite/dao"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/response"
	"tiktok-lite/util"
)

type UserController struct{}

// UserInfoOther 获取用户信息
// 注意：这是到达别人的主页，返回别人的信息
func UserInfoOther(context *gin.Context) {
	// 分别获取发布视频者id和用户token
	// 注意：这块使用context.query()获取的是字符串类型，需要将字符串转换为int64
	var publishId int64
	publishId, _ = util.String10Bit2Int64(context.Query("user_id"))
	// 这块就不用鉴权了，能够进入到这块说明中间件那块已经鉴权过了，只需要获取用户信息
	// var user model.User
	user, _ := context.Get(global.CtxUserKey)
	// 查询视频发布者信息
	publisher, _ := dao.UserInfoById(publishId)
	// 查询用户是否已经关注这个视频发布者
	exists, _ := dao.UserFollower(user.(model.User).Id, publishId)

	// 构建响应对象
	userResp := response.UserResponse{
		Response: response.Response{StatusCode: 200, StatusMsg: "成功"},
		User: response.User{
			User:     publisher,
			IsFollow: exists,
		},
	}
	context.JSON(http.StatusOK, userResp)
}

// UserInfo 获取当前登录用户的信息
func UserInfo(context *gin.Context) {
	// 通过全局 Key 获取当前用户
	// user, _ := context.Get(global.CtxUserKey)
	u := CtxUser(context)
	if u == DefUser {
		context.JSON(http.StatusForbidden, response.BaseInputError("Token invalid!!"))
		return
	}

	// 封装用户信息
	userResp := response.User{
		User:     *u,
		IsFollow: true, // 由于这里是用户在主页看到自己信息，所以是默认关注自己的
	}
	context.JSON(http.StatusOK, response.UserResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "成功"},
		User:     userResp,
	})
}

// UserRegister 用户注册
func UserRegister(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	// 插入数据
	userId, err := dao.UserRegister(username, password)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.BaseInputError("不允许重复的 username"))
		return
	}
	// 生成token
	user := model.User{Id: userId, UserName: username, Name: username}
	token, _ := util.GetToken(user)
	context.JSON(http.StatusOK, response.UserTokenSuccess(userId, token, "注册成功"))
}

// UserLogin 用户登录
func UserLogin(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	// 查询用户是否存在
	user, err := dao.UserLogin(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, "账号或密码错误")
		return
	}

	// strings.Compare(s1,s2), 0代表相等，1代表s1>s2,-1代表s1<s2
	if strings.Compare(password, user.PassWord) == 0 {
		// 获取token
		user.PassWord = ""
		token, _ := util.GetToken(user)

		context.JSON(http.StatusOK, response.UserTokenSuccess(user.Id, token, "登陆成功"))
		return
	}
	context.JSON(http.StatusBadRequest, response.BaseInputError("账号或密码错误"))
}
