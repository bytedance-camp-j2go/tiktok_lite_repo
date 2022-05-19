package driver

import (
	_ "github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/qiniu"
	"go.uber.org/zap"
)

func init() {
	InitALlDrivers()
	zap.L().Debug("all driver init")
}

func InitALlDrivers() {

}
