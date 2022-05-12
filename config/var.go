package config

import (
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

var (
	Conf Config
	// Log 支持 format 输出, 而 zap.L() 只能 结构化输出数据, 但性能更高
	Log  *zap.Logger
	Logf *zap.SugaredLogger
)
