package qiniu

import (
	"context"
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/base"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/qiniu/go-sdk/v7/storage"
)

func init() {
	// 注册七牛存储器到 base.driverMap
	base.RegisterDriver(&QNDriver{})
}

type QNDriver struct {
	// 组合 (约等于继承, 调用未重写的方法会报错)
	base.Model
}

func (d QNDriver) Name() string {
	return "QiNiu"
}

var defCtx = context.Background()

func (d QNDriver) Upload(file model.FileStream, account *model.DriverAccount) (string, error) {
	// 根据 qn-sdk 配置机房
	cfg := getCfg(account)
	// 保存到存储上的路径也是访问路径
	key := fmt.Sprintf("%s/%s", file.ParentPath, file.Name)
	token := signToken(account)

	// 上传对象
	uploader := storage.NewFormUploader(&cfg)
	res := storage.PutRet{}

	err := uploader.Put(defCtx, &res, token, key, file.File, file.Size, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", account.GetHost(), key), nil
}
