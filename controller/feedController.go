package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"tiktok-lite/dao"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/response"
	"tiktok-lite/util"
	"time"
)

const (
	// 参考 douyin web 版本, 每次固定返回 5 条视频。
	videoCnt = 5
)

// Feed 提取参数
// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个.
// latest_time query 说明：可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
func Feed(ctx *gin.Context) {
	timeStr := ctx.Query("latest_time")
	start := ParsingTimestampStr(timeStr)
	user := *CtxUser(ctx)
	// 根据时间戳, 返回 list
	feedProcess(ctx, start, user)

	// 可以做的：记录用户数据....

}

// 处理数据
func feedProcess(ctx *gin.Context, start time.Time, user model.User) {
	// 需要计算 start 在 set 中的排名, 决定使用二分查找
	// 找到比 start 大的第一个元素的排名, 然后返回 start + offset 个视频信息
	max := strconv.FormatInt(start.UnixMilli(), 10)
	rangeBy := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    max,
		Offset: 0,
		Count:  videoCnt,
	}

	videoIdList, err := util.ZSetRangeByScoreInt(global.VideoSeqSetKey, rangeBy)

	if len(videoIdList) == 0 && err != nil {
		zap.L().Debug("get video list err!!", zap.Int("len", len(videoIdList)), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	videos := videoProcess(videoIdList, user.Id)

	ctx.JSON(http.StatusOK, response.FeedResponse{
		StatusCode: 0,
		VideoList:  videos,
		NextTime:   calNextTime(videos),
	})

}

// videoProcess 获取 video 信息并封装用户信息，封装
func videoProcess(videoIds []int64, uid int64) []response.Video {
	videos, err := dao.VideoQueryList(videoIds)
	if err != nil {
		zap.L().Error("获取视频信息失败!!", zap.Error(err))
		return nil
	}

	res := make([]response.Video, 0, len(videos))

	// 不在 for 循环内声明变量
	var author response.User
	for _, video := range videos {
		author, err = response.NewUser(video.UserId, uid)
		if err != nil {
			zap.L().Error("get UserInfo", zap.Error(err))
			res = append(res, response.Video{Video: video})
			continue
		}
		res = append(res, response.Video{Video: video, Author: author})
	}
	return res
}

//  ====================== 时间解析 ======================

func calNextTime(videos []response.Video) int64 {
	n := len(videos)
	if n == 0 {
		return util.TimeNowInt64()
	}
	lastVideo := videos[n-1]
	return lastVideo.UpdatedAt.UnixMilli()
}

// ParsingTimestampStr 解析时间戳字符串
func ParsingTimestampStr(timeStr string) time.Time {
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return time.Now()
	}
	return time.Unix(timestamp, 0)
}
