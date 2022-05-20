package base

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
)

// Base 统一提供 Driver 行为的定义
// 由项目的需求做驱动器的行为定义
// 上传、删除( 真删除 ), 标记删除由数据库完成
// 存储数据列表
type Base struct{}

var (
	// read only
	driverMap = map[string]Driver{}
)

func RegisterDriver(driver Driver) {
	global.Logf.Infof("register driver: 「%s」", driver.Name())
	driverMap[driver.Name()] = driver
}

func GetDriver(key string) (d Driver, ok bool) {
	d, ok = driverMap[key]
	return
}

func GetDriverMap() map[string]Driver {
	return driverMap
}
