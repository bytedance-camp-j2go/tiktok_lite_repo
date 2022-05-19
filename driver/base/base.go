package base

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

// Base 统一提供 Driver 行为的定义
// 由项目的需求做驱动器的行为定义
// 上传、删除( 真删除 ), 标记删除由数据库完成
// 存储数据列表
type Base struct{}

// func (b Base) Config() DriverConfig {
// 	return DriverConfig{}
// }
//
// func (b Base) Items() []Item {
// 	return nil
// }

func (b Base) Save(account *model.DriverAccount, old *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Base) File(path string, account *model.DriverAccount) (*model.File, error) {
	return nil, ErrNotImplement
}

func (b Base) Files(path string, account *model.DriverAccount) ([]model.File, error) {
	return nil, ErrNotImplement
}

func (b Base) Link(args Args, account *model.DriverAccount) (*Link, error) {
	return nil, ErrNotImplement
}

func (b Base) Path(path string, account *model.DriverAccount) (*model.File, []model.File, error) {
	return nil, nil, ErrNotImplement
}

func (b Base) Preview(path string, account *model.DriverAccount) (interface{}, error) {
	return nil, ErrNotImplement
}

func (b Base) Move(src string, dst string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Base) Rename(src string, dst string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Base) Copy(src string, dst string, account *model.DriverAccount) error {
	return ErrNotImplement
}

func (b Base) Delete(path string, account *model.DriverAccount) error {
	return ErrNotImplement
}

// func (b Base) Upload(file *model.FileStream, account *model.DriverAccount) error {
// 	return ErrNotImplement
// }
