package base

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

/* todo
Config 配置
Config() DriverConfig
	- 抽象 Config 对象约束各种不同存储引擎的行为
	- 现在存储器少用不上
	- 会在考虑支持更多存储器

// type DriverConfig struct {
// 	Name     string
// 	HostName string
// }
*/

// Driver 存储驱动的接口定义
type Driver interface {
	Name() string
	// Save 保存时处理
	Save(account *model.DriverAccount, old *model.DriverAccount) error

	// File 取文件
	File(path string, account *model.DriverAccount) (*model.File, error)
	// Files 取文件夹
	Files(path string, account *model.DriverAccount) ([]model.File, error)

	// Move 移动/改名
	Move(src string, dst string, account *model.DriverAccount) error
	// Rename 改名
	Rename(src string, dst string, account *model.DriverAccount) error
	// Copy 拷贝
	Copy(src string, dst string, account *model.DriverAccount) error
	// Delete 删除
	Delete(path string, account *model.DriverAccount) error

	// Upload 上传
	Upload(file *model.FileStream, account *model.DriverAccount) (string, error)
}
