package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"gorm.io/gorm"
	"tiktok-lite/config"
)

var (
	// Conf Config、全局配置对象
	Conf config.Config

	// Log == zap.L() 结构化输出数据, 但性能更高
	Log *zap.Logger

	// Logf 支持 format 的日志
	Logf *zap.SugaredLogger

	// DB 获取mysql连接
	DB *gorm.DB

	// RedisDB Redis 连接
	RedisDB *redis.Client
)
