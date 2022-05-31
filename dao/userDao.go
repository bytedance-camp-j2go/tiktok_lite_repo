package dao

import (
	"errors"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"gorm.io/gorm"
)

// UserInfoById 根据userId获取用户信息
func UserInfoById(userId int64) (model.User, error) {
	db := global.DB
	user := model.User{}
	err := db.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// UserLogin 用户登录，查询用户是否存在
func UserLogin(username string) (model.User, error) {
	db := global.DB
	user := model.User{}
	err := db.Where("username = ?", username).Find(&user).Error
	// 说明为空
	if user == (model.User{}) {
		return user, err
	}
	return user, nil
}

// UserRegister 用户注册
func UserRegister(username string, password string) (int64, error) {
	db := global.DB
	// 先查询这个用户名是否存在
	user, err := UserLogin(username)
	if user != (model.User{}) {
		return -1, err
	}

	user = model.User{Id: util.UniqueID(), UserName: username, PassWord: password}
	err = db.Create(&user).Error
	if err != nil {
		return -1, err
	}
	return user.Id, nil
}

// UserFollower 查询user用户的关注列表中是否有publisher用户
func UserFollower(userId int64, publisherId int64) (bool, error) {
	db := global.DB
	userFollower := model.UserFollower{}
	err := db.Where("user_id = ? and follow_user_id = ?", userId, publisherId).Find(&userFollower).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {

	}
	if userFollower == (model.UserFollower{}) {
		return false, err
	}
	return true, nil
}
