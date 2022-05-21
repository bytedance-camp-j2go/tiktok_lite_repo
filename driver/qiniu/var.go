package qiniu

import (
	"time"
)

const (
	// token 缓存的超时时间
	expiresRedis  = 7200 * time.Second
	expiresOffset = 10 * time.Second
	// 向 Kodo 服务端申请 Token 时设置的 Token 过期时间
	expiresQBox = expiresRedis + expiresOffset

	ZoneHuanan  = "Huanan"
	ZoneHuabei  = "Huabei"
	ZoneHuadong = "Huadong"
)
