package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
	"testing"
	"tiktok-lite/bootstrap"
	_ "tiktok-lite/driver"
	"tiktok-lite/driver/operate"
	"tiktok-lite/global"
	"tiktok-lite/model"
	"tiktok-lite/util"
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

func testUpload(path, name string, size int64, reader io.Reader) {
	fileStream := model.FileStream{
		ParentPath: path,
		Name:       name,
		Size:       size,
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

func TestUploadJSON(t *testing.T) {
	bootstrap.InitAll()

	reader := strings.NewReader(`{"msg":"hello!"}`)
	user := model.User{
		Id:       123123,
		UserName: "",
		PassWord: "",
		Name:     "",
	}

	testUpload(
		fmt.Sprintf("%v/%s", user.Id, util.GetNowFormatTodayTime()),
		fmt.Sprintf("test-%d.json", time.Now().Unix()),
		reader.Size(),
		reader,
	)
}

func TestUploadLocalFile(t *testing.T) {
	bootstrap.InitAll()
	sysUser, _ := user.Current()
	filePath := sysUser.HomeDir
	fileName := "1942357430.mp4"
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", filePath, fileName), os.O_CREATE|os.O_APPEND, 6)
	if err != nil {
		return
	}

	info, err := file.Stat()

	testUpload(
		"/test-video", info.Name(),
		info.Size(),
		file,
	)
}
