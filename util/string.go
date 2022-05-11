package util

import (
	"time"
)

const (
	TimeFormatLayout = "2006-01-02|03:04:05"
	DayFormatLayout  = "2006-01-02"
)

func GetNowFormatTodayTime() string {
	return time.Now().Format(DayFormatLayout)
}
