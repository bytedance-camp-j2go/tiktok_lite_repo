package qiniu

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// TODO IMPL QinNiu SDK

const (
	maxRetry        = 3
	qnTokenCacheKey = "qiniu:token"
)

// 通过 SDK 签 TOKEN 存入 REDIS 使用事先 GET 没有则重新获取并存入
func getUploadToken(driver *model.DriverAccount) string {
	// 从缓存获得 Token
	exist, err := util.ExistKey(qnTokenCacheKey)
	if err != nil || !exist {
		return signToken(driver)
	}
	token, err := util.GetStringFromRedis(qnTokenCacheKey)
	// 应该是极小概率事件
	if err != nil {
		global.Logf.Errorf("get token error >> %q", err)
	}
	return token
}

func signToken(driver *model.DriverAccount) string {
	// 序列化到 redis? 可能不需要
	putPolicy := storage.PutPolicy{
		Scope:   driver.Bucket,
		Expires: expiresQBox,
		// MimeLimit 设置上传数据格式, 如果没有设置可能会根据文件名自动判断
		// MimeLimit: "!application/json;text/plain",
	}
	mac := qbox.NewMac(driver.AccessKey, driver.SecretKey)

	// 根据上传策略申请 token
	upToken := putPolicy.UploadToken(mac)

	go func() {
		util.Save2Redis("", []byte(upToken), expiresRedis)
	}()
	return upToken
}
