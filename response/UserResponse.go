/**
* @Author:drl
* @Date: 2022/5/19 13:02
 */
package response

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

//返回用户信息
type UserResponse struct {
	Response
	User model.User `json:"user"` //用户信息
}

//用户登录成功返回响应
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"` //用户id
	Token  string `json:"token"`   //鉴权token
}

//用户注册成功返回响应
type UserRegisterResponse struct {
	Response Response
	UserId   int64  `json:"user_id"` //用户id
	Token    string `json:"token"`   //鉴权token
}
