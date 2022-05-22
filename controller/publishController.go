package controller

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/operate"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func PublishAction(context *gin.Context) {
	//获取用户信息
	var a any
	a, _ = context.Get(global.UserName)
	user := a.(model.User)
	//获取视频信息
	data, err := context.FormFile("data")
	if err != nil {
		context.JSON(http.StatusOK, response.PublishActionResponse{
			//注意：状态码 0成功，其他失败
			Response: response.Response{StatusCode: 1, StatusMsg: "上传失败"}},
		)
	}
	//获取视频文件名称，只是视频文件名称及后缀，
	//例：/path/user/test.txt ---> test.txt
	finalName := filepath.Base(data.Filename)
	//创建fileStream流
	file, _ := data.Open()
	fileStream := model.FileStream{
		File:       file,
		Size:       int64(data.Size),
		ParentPath: "/test/",
		Name:       finalName,
		MIMEType:   context.ContentType(),
	}
	driverAccount := model.GetDriverAccount(fileStream.ParentPath)
	//上传文件，res是上传之后的视频url
	videoUrl, err := operate.Upload(&driverAccount, fileStream)
	//获取视频封面url
	picUrl := videoUrl + "?vframe/jpg/offset/1"
	if err != nil {
		//注意：状态码 0成功 其他值失败
		context.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	dao.PublishActionDao(user, videoUrl, picUrl)
	context.JSON(http.StatusOK, response.Response{StatusCode: 0, StatusMsg: picUrl})
}
