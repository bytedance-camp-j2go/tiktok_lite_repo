/**
* @Author:drl
* @Date: 2022/5/19 0:52
 */
package dao

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

//根据userId获取用户信息
func GetUserInfo() {
	//这里可以直接通过config.DB来获取数据库连接
	//在bootstrap/db.go中已经进行初始化
	db := config.DB
	user := model.User{}
	//fmt.Println(db)
	db.Take(&user)
	fmt.Println(user)
}
