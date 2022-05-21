package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	_ "github.com/bytedance-camp-j2go/tiktok_lite_repo/driver"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/operate"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"strings"
	"testing"
	"time"
)

// web 表单测试上传示例代码, 供使用者参考
/*
	参考 https://github.dev/Xhofe/alist
	form, err := ctx.MultipartForm()
	if err != nil {
		common.ErrorResp(c, err, 400)
	}
	files := form.File["files"]
	if err != nil {
		return
	}
	for i, file := range files {
		open, err := file.Open()
		fileStream := model.FileStream{
			File:       open,
			Size:       uint64(file.Size),
			ParentPath: path_,
			Name:       file.Filename,
			MIMEType:   file.Header.Get("Content-Type"),
		}
		clearCache := false
		if i == len(files)-1 {
			clearCache = true
		}
		err = operate.Upload(driver, account, &fileStream, clearCache)
		if err != nil {
			if i != 0 {
				_ = base.DeleteCache(path_, account)
			}
			common.ErrorResp(c, err, 500)
			return
		}
	}

*/

func TestUpload(t *testing.T) {
	bootstrap.InitAll()

	reader := strings.NewReader(`{"msg":"hello!"}`)
	user := model.User{
		Id:       123123,
		UserName: "",
		PassWord: "",
		Name:     "",
	}
	fileStream := model.FileStream{
		ParentPath: fmt.Sprintf("%v/%s", user.Id, util.GetNowFormatTodayTime()),
		Name:       fmt.Sprintf("test-%d.json", time.Now().Unix()),
		Size:       reader.Size(),
		File:       reader,
	}

	account := model.GetDriverAccount(fileStream.ParentPath)
	global.Logf.Infof("use account >> %v", account)
	url, err := operate.Upload(&account, fileStream)
	if err != nil {
		global.Logf.Errorf("upload err!! %v\n", err)
	} else {
		global.Logf.Infof("visit >> %v\n", url)
	}

}
