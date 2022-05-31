package bootstrap

import (
	"github.com/go-redis/redis/v8"
	"tiktok-lite/global"
)

func InitRedis() {

	global.RedisDB = redis.NewClient(&redis.Options{
		Addr:     global.Conf.Redis.SVN(),
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.No,
	})

	if err := global.RedisDB.Ping(global.RedisDB.Context()).Err(); err != nil {
		panic("Redis 连接失败 >> " + err.Error())
	}
}
