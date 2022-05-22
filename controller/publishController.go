package controller

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/operate"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// PublishAction 视频投稿
func PublishAction(context *gin.Context) {
	// 获取用户信息
	var a any
	a, _ = context.Get(global.UserName)
	user := a.(model.User)
	// 获取视频信息
	data, err := context.FormFile("data")
	if err != nil {
		context.JSON(http.StatusOK, response.PublishActionResponse{
			// 注意：状态码 0成功，其他失败
			Response: response.Response{StatusCode: 1, StatusMsg: "上传失败"}},
		)
	}
	// 获取视频文件名称，只是视频文件名称及后缀，
	// 例：/path/user/test.txt ---> test.txt
	finalName := filepath.Base(data.Filename)
	// 创建fileStream流
	file, _ := data.Open()
	fileStream := model.FileStream{
		File:       file,
		Size:       int64(data.Size),
		ParentPath: "/test/",
		Name:       finalName,
		MIMEType:   context.ContentType(),
	}
	driverAccount := model.GetDriverAccount(fileStream.ParentPath)
	// 上传文件，res是上传之后的视频url
	playUrl, err := operate.Upload(&driverAccount, fileStream)
	// 获取视频封面url
	coverUrl := playUrl + "?vframe/jpg/offset/1"
	// 获取视频标题,，这块使用query()
	title := context.Query("title")
	if err != nil {
		// 注意：状态码 0成功 其他值失败
		context.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	err = dao.PublishActionDao(user, playUrl, coverUrl, title)
	if err != nil {
		// 注意：状态码 0成功 其他值失败
		context.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	context.JSON(http.StatusOK, response.Response{StatusCode: 0, StatusMsg: coverUrl})
}

// PublishList 发布列表
// 场景：即当用户点开某一视频up主页，可以看到其视频列表
func PublishList(context *gin.Context) {
	// 获取up主信息
	var a any
	a = context.Query("user_id")
	publisherId := a.(int64)
	publisher, _ := dao.UserInfoById(publisherId)
	// 查询up主发布的视频列表
	fmt.Println(publisher)

}
