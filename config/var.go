package config

import (
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Conf Config、全局配置对象
	Conf Config

	// Log == zap.L() 结构化输出数据, 但性能更高
	Log *zap.Logger

	// Logf 支持 format 的日志
	Logf *zap.SugaredLogger

	// 获取mysql连接
	DB *gorm.DB
)
