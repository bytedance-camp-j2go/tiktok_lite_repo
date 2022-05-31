package bootstrap

import (
	"tiktok-lite/global"
	"tiktok-lite/model"
)

func InitModel() {
	err := global.DB.AutoMigrate(
		&model.DriverAccount{},
		&model.User{},
		&model.UserFavorite{},
		&model.UserFollower{},
		&model.Video{},
	)

	if err != nil {
		global.Logf.Errorf("driver account init error! | %v\n", err)
	}
}
