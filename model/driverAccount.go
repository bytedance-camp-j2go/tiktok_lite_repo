package model

import (
	"gorm.io/gorm"
)

type DriverAccount struct {
	gorm.Model
	Name string `json:"name" gorm:"unique" binding:"required"` // 存储账户的唯一名称
	Type string `json:"type"`                                  // 类型，即 driver 名

	// for QiNiu
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Host      string `json:"host"`
	UseHTTPS  bool   `json:"use_https"`
	// 机房位置
	Zone string `json:"zone"`
}
