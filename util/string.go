package util

import (
	"strconv"
	"time"
)

const (
	TimeFormatLayout = "2006-01-02|03:04:05"
	DayFormatLayout  = "2006-01-02"
)

func GetNowFormatTodayTime() string {
	return time.Now().Format(DayFormatLayout)
}

func String10Bit2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func Int64D2String(i int64) string {
	return strconv.FormatInt(i, 10)
}
