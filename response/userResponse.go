package response

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

// User 返回用户信息中的User对象
type User struct {
	// Id            int64  `json:"id"`             // 视频发布者id
	// Name          string `json:"name"`           // 视频发布者昵称
	// FollowCount   int64  `json:"follow_count"`   // 视频发布者关注数量
	// FollowerCount int64  `json:"follower_count"` // 视频发布者粉丝数量
	model.User
	IsFollow bool `json:"is_follow"` // 用户是否关注这个视频发布者

	// 随机图片 API
	// Just for fun
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	// 随机图 API 用作背景 URL
	// https://picsum.photos/400/200?grayscale
	// # 随机头像API请求方式 #
	// Method: GET
	// # 请求地址 #
	// https://api.sunweihu.com/api/sjtx/api.php
	// # 参数 #
	// lx【1.男头：a1 2.女头：b1 3.动漫：c1 4.动漫女头：c2 5.动漫男头：c3】
	// # 返回数据 #
	// 本API无返回数据,直接输出头像
	// # 备注 #
	// 后续将会增加更多类型的头像，欢迎大家投稿
	// # 示例 #
	// https://api.sunweihu.com/api/sjtx/api.php?lx=c1
}

const (
	avatarAPI     = "https://api.sunweihu.com/api/sjtx/api.php"
	backGroundAPI = "https://picsum.photos/400/200?grayscale"
)

// NewUser 查询 User 并计算是否已关注
// uid = 被查询用户 id, u2id = 查询发起者 id
func NewUser(uid, u2id int64) (User, error) {
	if id, err := dao.UserInfoById(uid); err == nil {
		return User{
			id,
			isFollow(uid, u2id),
			avatarAPI,
			"测试签名",
			backGroundAPI,
		}, err
	}

	return UserError, nil
}

func isFollow(uid, u2id int64) bool {
	if u2id == -1 {
		return false
	}
	// todo
	return false
}

// UserResponse 返回用户信息
type UserResponse struct {
	Response
	User User `json:"user"`
}

// UserTokenResponse 用户登录成功返回响应
// 登陆、注册成功都是返回 uid + token
// 索性合并两种 Resp，
// 可以更方便的封装返回方法
type UserTokenResponse struct {
	Response
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 鉴权token
}

func UserTokenSuccess(uid int64, token, msg string) UserTokenResponse {
	return UserTokenResponse{BaseSuccess(msg), uid, token}
}

// // UserRegisterResponse 用户注册成功返回响应
// type UserRegisterResponse struct {
// 	Response
// 	UserId int64  `json:"user_id"` // 用户id
// 	Token  string `json:"token"`   // 鉴权token
// }
//
// func UserRegisterSuccess(uid int64, token, msg string) UserRegisterResponse {
// 	return UserRegisterResponse{BaseSuccess(msg), uid, token}
// }
