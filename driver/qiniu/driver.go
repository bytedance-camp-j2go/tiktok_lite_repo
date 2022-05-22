package qiniu

import (
	"context"
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/base"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

func Init() {
	// 注册七牛存储器到 base.driverMap
	base.RegisterDriver(&QNDriver{})
}

type QNDriver struct {
	// 组合 (约等于继承, 调用未重写的方法会报错)
	base.Model
}

func (d QNDriver) Name() string {
	return "qiniu"
}

var (
	defCtx = context.Background()
)

const previewParams = "vframe/jpg/offset/1"

func (d QNDriver) Upload(file model.FileStream, account *model.DriverAccount) (string, error) {
	// 根据 qn-sdk 配置机房
	cfg := getCfg(account)
	// 保存到存储上的路径也是访问路径
	key := fmt.Sprintf(
		"%s/%s",
		util.ParentPath2Suffix(file.ParentPath),
		file.Name,
	)
	token := getUploadToken(account)

	// 上传对象
	uploader := storage.NewFormUploader(&cfg)
	res := storage.PutRet{}

	err := uploader.Put(defCtx, &res, token, key, file.File, file.Size, nil)
	if err != nil {
		return "", err
	}

	zap.L().Info("upload success!",
		zap.String("file-key", res.Key),
		zap.String("file-hash", res.Hash),
		zap.String("ps-id", res.PersistentID),

		zap.String("driver-name", d.Name()),
		zap.String("account", account.Name),
	)

	return fmt.Sprintf("%s/%s", account.GetHost(), key), nil
}

// Preview 通过视频的访问连接处理视频预览链接 (提供默认的视频封面
func (d QNDriver) Preview(url string) (string, error) {
	return fmt.Sprintf("%s?%s", url, previewParams), nil
}
