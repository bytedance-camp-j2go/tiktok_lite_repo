package base

import (
	"tiktok-lite/model"
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
	// Name 返回 driver 的命名 == Account 的type, 全局唯一
	Name() string

	// Host 返回驱动器内自定义的 host url, 直接或得驱动器对象, 从而支持自行拼接 url
	// 注意!! 实现 Host() 是可选, 如果没有实现就进行调用会报错的.
	// Host() (string, error)
	// 经过考虑之后觉得不需要 driver 返回 host, 因为一个 driver 对象需要涉及多个账户

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
	Upload(file model.FileStream, account *model.DriverAccount) (string, error)

	Preview(url string) (string, error)
}
