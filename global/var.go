package global

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

var (
	// Conf Config、全局配置对象
	Conf config.Config

	// Log == zap.L() 结构化输出数据, 但性能更高
	Log *zap.Logger

	// Logf 支持 format 的日志
	Logf *zap.SugaredLogger

	// RedisDB Redis 连接
	RedisDB *redis.Client
	// Mysql 连接

)
