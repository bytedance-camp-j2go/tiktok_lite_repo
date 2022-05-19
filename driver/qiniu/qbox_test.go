package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gorm.io/gorm"
	"testing"
)

var (
	driver = model.DriverAccount{
		Model:     gorm.Model{},
		Name:      "qbox-01",
		Type:      "qbox",
		AccessKey: "O_qaFxMC9xvkbMURCs5Dhcxf1EDZnUJKIlry72rh",
		SecretKey: "v7YJbHOh-zzKQPw9o-7D60uxSpHdpSObCL7ZOcMF",
		Bucket:    "j2go",
		Host:      "dy-resource.evlic.cn",
		UseHTTPS:  true,
		Zone:      "Huanan",
	}

	cfg = storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      true,
		UseCdnDomains: false,
	}
)

func TestUpload(t *testing.T) {
	putPolicy := storage.PutPolicy{
		Scope:     driver.Bucket,
		MimeLimit: "!application/json;text/plain",
	}

	mac := qbox.NewMac(driver.AccessKey, driver.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	saveName := "user/path/to/file/xx1"
	data := []byte(`{"msg":"hello, this upload Demo By j2go", "for": "test path/to/file"}`)
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, saveName, bytes.NewReader(data), dataLen, &putExtra)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	saveName = "user/path/to/file/xx2"
	err = formUploader.Put(context.Background(), &ret, upToken, saveName, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
