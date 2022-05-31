package driver

import (
	"go.uber.org/zap"
	"tiktok-lite/driver/qiniu"
)

func Init() {
	InitALlDrivers()
	zap.L().Debug("all driver init")
}

// InitALlDrivers future-feature 通过这个方法初始化 config
func InitALlDrivers() {
	qiniu.Init()
}
