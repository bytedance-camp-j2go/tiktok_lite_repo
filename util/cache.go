package util

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
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
	if err != nil || err == redis.Nil {
		return false, err
	}
	return cnt > 0, nil
}

func GetStringFromRedis(key string) (string, error) {
	res, err := global.RedisDB.Get(global.RedisDB.Context(), key).Result()
	if err != nil {
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

// ZSetRangeByScoreStrings 范围查找值
func ZSetRangeByScoreStrings(key string, z *redis.ZRangeBy) ([]string, error) {
	return global.RedisDB.ZRangeByScore(global.RedisDB.Context(), key, z).Result()
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
// 	// global.RedisDB.
// 	args := []interface{}{"zrangebyscore", key, z.Min, z.Max}
// 	if z.Offset != 0 || z.Count != 0 {
// 		args = append(args, "limit", z.Offset, z.Count)
// 	}
//
// 	// 一下仿照: global.RedisDB.Do()
// 	ctx := global.RedisDB.Context()
// 	cmd := redis.NewIntSliceCmd(ctx, args...)
// 	// cmd := redis.NewStringSliceCmd(ctx, args...)
// 	_ = global.RedisDB.Process(ctx, cmd)
// 	return cmd.Result()
// }

// ZSetRangeByScoreInt 执行范围查询 如果 []int64 为空, 则 error 不可忽略
func ZSetRangeByScoreInt(key string, z *redis.ZRangeBy) ([]int64, error) {
	strings, err := ZSetRangeByScoreStrings(key, z)
	if err != nil {
		return nil, err
	}

	res := make([]int64, 0, len(strings))
	var p int64
	for idx := range strings {
		p, err = strconv.ParseInt(strings[idx], 64, 10)
		if err != nil {
			return res, err
		}
		res = append(res, p)
	}
	return res, nil
}

func ZSetCnt(key string) (int64, error) {
	return global.RedisDB.ZCard(global.RedisDB.Context(), key).Result()
}
