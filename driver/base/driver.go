package base

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"net/http"
)

type DriverConfig struct {
	Name          string
	OnlyProxy     bool // 必须使用代理（本机或者其他机器）
	OnlyLocal     bool // 必须本机返回的
	ApiProxy      bool // 使用API中转的
	NoNeedSetLink bool // 不需要设置链接的
	NoCors        bool // 不可以跨域
	LocalSort     bool // 本地排序
}

type Item struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Values      string `json:"values"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type Args struct {
	Path   string
	IP     string
	Header http.Header
}

// Driver 存储驱动的接口定义
type Driver interface {
	// Config 配置
	Config() DriverConfig
	// Items 初始化驱动器所需参数
	Items() []Item
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
	// Upload(file *model.FileStream, account *model.DriverAccount) error
}
