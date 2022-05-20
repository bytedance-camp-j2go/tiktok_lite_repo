package model

import (
	"gorm.io/gorm"
)

// DriverAccount 通过接口向数据库存入驱动器配置
type DriverAccount struct {
	gorm.Model
	Name string `json:"name" gorm:"unique" binding:"required"` // 存储账户的唯一名称
	Type string `json:"type" gorm:"type"`                      // 类型，即 driver 名

	// for QiNiu
	AccessKey string `json:"access_key" gorm:"column:qn_ak"`
	SecretKey string `json:"secret_key" gorm:"column:qn_sk"`
	Bucket    string `json:"bucket"`
	Host      string `json:"host"`
	UseHTTPS  bool   `json:"use_https"`
	// 机房位置
	Zone string `json:"zone"`
}

// GetDriverAccount todo-feature 通过 path 匹配存储账户
// 现行逻辑直接返回七牛的存储账户
// 上传时调用
func GetDriverAccount(path string) DriverAccount {
	return DriverAccount{}
}
