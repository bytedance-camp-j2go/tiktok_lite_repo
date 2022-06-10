package util

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"tiktok-lite/global"
	"time"
)

// 封装 Redis 增删、简单操作

var (
	redisClient *redis.Client
	redisDefCtx context.Context
)

func InitRedis() {
	redisClient = global.RedisDB
	redisDefCtx = redisClient.Context()
}

func Save2Redis(key string, v []byte, expires time.Duration) {
	result, err := redisClient.Set(redisDefCtx, key, v, expires).Result()
	if err != nil {
		zap.L().Error("写入 redis 失败", zap.Error(err))
		return
	}
	zap.L().Debug("save2redis", zap.String("say", result))
}

func ExistKey(key string) (bool, error) {
	cnt, err := redisClient.Exists(redisDefCtx, key).Result()
	if err != nil || err == redis.Nil {
		return false, err
	}
	return cnt > 0, nil
}

func GetStringFromRedis(key string) (string, error) {
	res, err := redisClient.Get(redisDefCtx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func ZAdd2Redis(key string, score float64, v any) {
	redisClient.ZAdd(
		redisDefCtx,
		key,
		&redis.Z{
			Score: score, Member: v,
		},
	)
}

func ZRM2Redis(key string, ms ...any) int64 {
	cnt, err := redisClient.ZRem(
		redisDefCtx,
		key,
		ms,
	).Result()
	if err != nil {
		zap.L().Error("redis del error!!", zap.String("Key", key), zap.Error(err))
	}
	return cnt
}

// ZSetRevRangeByScoreStrings 范围查找值
func ZSetRevRangeByScoreStrings(key string, z *redis.ZRangeBy) ([]string, error) {
	// return redisClient.ZRangeByScore(redisDefCtx, key, z).Result()
	return redisClient.ZRevRangeByScore(redisDefCtx, key, z).Result()
}

// ZSetRangeByScoreInt
// 写法参照两段源码 Do ZRangeByScore
/*
 	args := []interface{}{zcmd, key, opt.Min, opt.Max}
	if withScores {
		args = append(args, "withscores")
	}
	if opt.Offset != 0 || opt.Count != 0 {
		args = append(
			args,
			"limit",
			opt.Offset,
			opt.Count,
		)
	}
	cmd := NewStringSliceCmd(ctx, args...)
	_ = c(ctx, cmd)
	return cmd
*/

/**/
//
// func ZSetRangeByScoreInt(key string, z *redis.ZRangeBy) ([]int64, error) {
// 	// func ZSetRangeByScoreInt(key string, z *redis.ZRangeBy) ([]string, error) {
// 	// redisClient.
// 	args := []interface{}{"zrangebyscore", key, z.Min, z.Max}
// 	if z.Offset != 0 || z.Count != 0 {
// 		args = append(args, "limit", z.Offset, z.Count)
// 	}
//
// 	// 一下仿照: redisClient.Do()
// 	ctx := redisDefCtx
// 	cmd := redis.NewIntSliceCmd(ctx, args...)
// 	// cmd := redis.NewStringSliceCmd(ctx, args...)
// 	_ = redisClient.Process(ctx, cmd)
// 	return cmd.Result()
// }

// ZSetRangeByScoreInt 执行范围查询 如果 []int64 为空, 则 error 不可忽略
func ZSetRangeByScoreInt(key string, z *redis.ZRangeBy) ([]int64, error) {
	strings, err := ZSetRevRangeByScoreStrings(key, z)
	if err != nil {
		return nil, err
	}

	res := make([]int64, 0, len(strings))
	var p int64
	for idx := range strings {
		p, err = String10Bit2Int64(strings[idx])
		if err != nil {
			return res, err
		}
		res = append(res, p)
	}
	return res, nil
}

func ZSetCnt(key string) int64 {
	result, err := redisClient.ZCard(redisDefCtx, key).Result()
	if err != nil {
		zap.L().Error("redis error!!", zap.String("Key", key), zap.Error(err))
	}
	return result
}

// ZSetRank 返回 m 在 redis 中的排名，如果不存在返回 -1
func ZSetRank(key, m string) int64 {
	res, err := redisClient.ZRank(redisDefCtx, key, m).Result()
	if err != nil {
		// err == redis.Nil
		// zap.L().Debug("ZSetRank Error!!", zap.Error(err), zap.String("z-key", key), zap.String("z-m", m))
		return -1
	}
	return res
}
