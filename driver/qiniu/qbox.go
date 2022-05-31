package qiniu

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/util"
)

// TODO IMPL QinNiu SDK

const (
	maxRetry        = 3
	qnTokenCacheKey = "qiniu:token:%s"
)

// 通过 SDK 签 TOKEN 存入 REDIS 使用事先 GET 没有则重新获取并存入
func getUploadToken(account *model.DriverAccount) string {
	redisTokenKey := fmt.Sprintf(qnTokenCacheKey, account.Name)
	// 从缓存获得 Token
	exist, err := util.ExistKey(redisTokenKey)
	if err != nil || !exist {
		zap.L().Debug("get token err!! will reSign", zap.Error(err))
		return signToken(account)
	}
	token, err := util.GetStringFromRedis(redisTokenKey)
	// 应该是极小概率事件
	if err != nil {
		global.Logf.Errorf("get token error >> %q\n", err)
	}
	zap.L().Debug("token >> ", zap.String("cache", token))
	return token
}

// 获取上传配置, 主要是存储器的机房区域
func getCfg(account *model.DriverAccount) storage.Config {
	var zone *storage.Region

	switch account.Zone {
	case ZoneHuanan:
		zone = &storage.ZoneHuanan
	case ZoneHuadong:
		zone = &storage.ZoneHuadong
	case ZoneHuabei:
		zone = &storage.ZoneHuabei
	default:
		// TODO 补全支持
		zap.L().Error("暂不支持此区域")
		panic("no support zone")
	}

	return storage.Config{
		Zone:          zone,
		UseHTTPS:      true,
		UseCdnDomains: false,
	}

}

// 签 token
func signToken(account *model.DriverAccount) string {
	// 序列化到 redis? 可能不需要
	putPolicy := storage.PutPolicy{
		Scope:   account.Bucket,
		Expires: uint64(expiresQBox.Seconds()),
		// MimeLimit 设置上传数据格式, 如果没有设置可能会根据文件名自动判断
		// MimeLimit: "!application/json;text/plain",
	}
	mac := qbox.NewMac(account.AccessKey, account.SecretKey)

	// 根据上传策略申请 token
	upToken := putPolicy.UploadToken(mac)

	go func() {
		util.Save2Redis(
			fmt.Sprintf(qnTokenCacheKey, account.Name), []byte(upToken), expiresRedis)
	}()
	return upToken
}
