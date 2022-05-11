package config

import (
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

var (
	Conf Config
	Log  *zap.Logger
)
