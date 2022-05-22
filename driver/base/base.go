package base

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

// Model 统一提供 Driver 行为的定义
// 由项目的需求做驱动器的行为定义
// 上传、删除( 真删除 ), 标记删除由数据库完成
// 存储数据列表
type Model struct{}

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

// 以下为未实现的接口方法. 除了 Name 以外都是不能调用的。

func (b Model) Name() string {
	return "Base"
}

func (b Model) Save(account *model.DriverAccount, old *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Model) File(path string, account *model.DriverAccount) (*model.File, error) {
	return nil, ErrNotImplement
}

func (b Model) Files(path string, account *model.DriverAccount) ([]model.File, error) {
	return nil, ErrNotImplement
}

func (b Model) Move(src string, dst string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Model) Rename(src string, dst string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Model) Copy(src string, dst string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Model) Delete(path string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Model) Upload(file model.FileStream, account *model.DriverAccount) (string, error) {
	return "", ErrNotImplement
}

func (b Model) Preview(url string) (string, error) {
	return "", ErrNotImplement
}
