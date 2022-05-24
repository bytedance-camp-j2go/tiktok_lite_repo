package util

import (
	"github.com/godruoyi/go-snowflake"
)

func UniqueID() uint64 {
	return snowflake.ID()
}
