/**
* @Author:drl
* @Date: 2022/5/19 0:48
 */
package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserController struct{}

//获取用户信息
func (UserController) User(context *gin.Context) {

}

//用户注册

func (UserController) Register(context *gin.Context) {
	//获取注册账号密码

}

//用户登录
func (UserController) Login(context *gin.Context) {
	//获取账号密码
	username := context.Query("username")
	password := context.Query("password")
	//查询用户是否存在
	user, err := dao.UserLogin(username, password)
	if !err {
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
