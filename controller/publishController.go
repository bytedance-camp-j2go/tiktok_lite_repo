package controller

import (
	"fmt"
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
	fmt.Println(finalName, user)
	//fileStream := model.FileStream{
	//	File:       model.FileStream{},
	//	Size:       data.Size,
	//	ParentPath: "/",
	//	Name:       finalName,
	//	MIMEType:   context.ContentType(),
	//}

}
