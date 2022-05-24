package util

import (
	"time"
)

// TimeNowInt64 返回当前时间的秒数
func TimeNowInt64() int64 {
	return time.Now().UnixMilli()
}
