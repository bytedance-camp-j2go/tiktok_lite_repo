package util

import (
	"github.com/godruoyi/go-snowflake"
)

func UniqueID() int64 {
	return int64(snowflake.ID())
}
