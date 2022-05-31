package main

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"testing"
	"tiktok-lite/bootstrap"
	"tiktok-lite/global"
	"tiktok-lite/util"
	"time"
)

//
func TestInsertVideoCache(t *testing.T) {
	bootstrap.InitAll()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		util.ZAdd2Redis(
			"test:"+global.VideoSeqSetKey,
			float64(util.TimeNowInt64()),
			i,
		)
	}

}

func TestZSetCnt(t *testing.T) {
	bootstrap.InitAll()
	cnt, err := util.ZSetCnt("test:" + global.VideoSeqSetKey)
	if err != nil {
		return
	}
	zap.L().Info("ZSet Cnt", zap.Int64("cnt", cnt))
}

func TestZSetRangeByScore(t *testing.T) {
	bootstrap.InitAll()
	res, err := util.ZSetRangeByScoreStrings("test:"+global.VideoSeqSetKey, &redis.ZRangeBy{
		Max:    strconv.FormatInt(util.TimeNowInt64(), 10),
		Min:    "-inf",
		Offset: 0,
		Count:  5,
	})
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	global.Logf.Infof("ZSetRange >> %#v\n", res)
}

func TestZSetRangeByScoreInt(t *testing.T) {
	bootstrap.InitAll()
	res, err := util.ZSetRangeByScoreInt("test:"+global.VideoSeqSetKey, &redis.ZRangeBy{
		Max:    strconv.FormatInt(util.TimeNowInt64(), 10),
		Min:    "-inf",
		Offset: 0,
		Count:  5,
	})
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	global.Logf.Infof("ZSetRange >> %#v\n", res)
}
