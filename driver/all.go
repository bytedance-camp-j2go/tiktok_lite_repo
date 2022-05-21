package driver

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/qiniu"
	"go.uber.org/zap"
)

func Init() {
	InitALlDrivers()
	zap.L().Debug("all driver init")
}

// InitALlDrivers future-feature 通过这个方法初始化 config
func InitALlDrivers() {
	qiniu.Init()
}
