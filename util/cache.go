package util

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

// 封装 Redis 增删、简单操作

func Save2Redis(key string, v []byte, expires time.Duration) {
	result, err := global.RedisDB.Set(global.RedisDB.Context(), key, v, expires).Result()
	if err != nil {
		zap.L().Error("写入 redis 失败", zap.Error(err))
		return
	}
	zap.L().Debug("save2redis", zap.String("say", result))

}

func ExistKey(key string) (bool, error) {
	cnt, err := global.RedisDB.Exists(global.RedisDB.Context(), key).Result()
	if err != redis.Nil {
		return false, err
	}
	return cnt > 0, nil
}

func GetStringFromRedis(key string) (string, error) {
	res, err := global.RedisDB.Get(global.RedisDB.Context(), key).Result()
	if err != redis.Nil {
		return "", err
	}
	return res, nil
}

func ZAdd2Redis(key string, score float64, v any) {
	global.RedisDB.ZAdd(
		global.RedisDB.Context(),
		key,
		&redis.Z{
			Score: score, Member: v,
		},
	)
}

// ZSetRangeByScore 范围查找值
func ZSetRangeByScore(key string, z *redis.ZRangeBy) ([]string, error) {

	return global.RedisDB.ZRangeByScore(global.RedisDB.Context(), key, z).Result()
}

func ZSetCnt(key string) (int64, error) {
	return global.RedisDB.ZCard(global.RedisDB.Context(), key).Result()
}
