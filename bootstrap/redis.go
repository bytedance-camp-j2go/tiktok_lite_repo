package bootstrap

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {

	global.RedisDB = redis.NewClient(&redis.Options{
		Addr:     global.Conf.Redis.String(),
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.No,
	})

	if err := global.RedisDB.Ping(global.RedisDB.Context()).Err(); err != nil {
		panic("Redis 连接失败 >> " + err.Error())
	}
}
