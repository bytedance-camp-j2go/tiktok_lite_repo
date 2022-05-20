/**
* @Author:drl
* @Date: 2022/5/19 0:52
 */
package dao

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

//根据userId获取用户信息
func GetUserInfo() {
	//这里可以直接通过config.DB来获取数据库连接
	//在bootstrap/db.go中已经进行初始化
	db := global.DB
	//user := model.User{}
	var user model.User
	//fmt.Println(db)
	db.Take(&user)
	fmt.Println(user)
}

//用户登录，查询用户是否存在
func UserLogin(username string, password string) (model.User, bool) {
	db := global.DB
	var user model.User
	db.Where("username=?", username).Find(&user)
	//说明为空
	if user == (model.User{}) {
		return user, false
	}
	fmt.Println("------------------->", user)
	return user, true
}
