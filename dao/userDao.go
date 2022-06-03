package dao

import (
	"errors"
	"gorm.io/gorm"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/util"
)

/*const (
	defRedisUserInfoCahceExpires = time.Second
	userCacheKey                 = "user:info:c:%d"
)*/

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

	user = model.User{Id: util.UniqueID(), UserName: username, Name: username, PassWord: password}
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

// // 通过userId和publisherId查询publisher信息
// func getUserById(userId int64, publisherId int64) (response.User, error) {
// 	db := global.DB
// 	var publisher model.User
// 	var publisherResp response.User
// 	err := db.Where("user_id=?", userId).Find(&publisher).Error
// 	if publisher == (model.User{}) {
// 		return publisherResp, err
// 	}
// 	isFollower, _ := UserFollower(userId, publisherId)
// 	// 拼接publisherResp
// 	publisherResp.User = publisher
// 	publisherResp.IsFollow = isFollower
// 	return publisherResp, nil
// }
