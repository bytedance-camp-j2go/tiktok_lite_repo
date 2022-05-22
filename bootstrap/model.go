package bootstrap

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

func InitModel() {
	err := global.DB.AutoMigrate(&model.DriverAccount{})

	if err != nil {
		global.Logf.Errorf("driver account init error! | %v\n", err)
	}
}
