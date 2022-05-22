package controller

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/dao"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/operate"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/response"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"path"
	"strings"
	"time"
)

func PublishAction(context *gin.Context) {
	// 从上下文获取用户信息
	ctxVal, _ := context.Get(global.UserName)
	user, ok := ctxVal.(model.User)
	if !ok {
		zap.L().Error("user info err!")
		context.JSON(http.StatusBadRequest, response.Response{StatusCode: 2, StatusMsg: ""})
		return
	}

	// 获取视频信息
	data, err := context.FormFile("data")
	if err != nil {
		context.JSON(http.StatusOK, response.PublishActionResponse{
			// 注意：状态码 0成功，其他失败
			Response: response.Response{StatusCode: 1, StatusMsg: "上传失败"}},
		)
	}

	// 获取视频文件名称，只是视频文件名称及后缀，
	// 例：test.txt ---> test-14:04:05.1231.txt
	finalName := fileName(data.Filename)

	// 创建fileStream流
	file, _ := data.Open()
	fileStream := model.FileStream{
		File: file,
		Size: data.Size,
		// 格式: userid/time
		ParentPath: fmt.Sprintf("/test/%b/%s", user.Id, util.GetNowFormatTodayTime()),
		Name:       finalName,
		MIMEType:   context.ContentType(),
	}

	driverAccount := model.GetDriverAccount(fileStream.ParentPath)
	// 上传文件，res是上传之后的视频url
	videoUrl, err := operate.Upload(&driverAccount, fileStream)
	// 获取视频封面url
	picUrl := videoUrl + "?vframe/jpg/offset/1"
	if err != nil {
		// 注意：状态码 0成功 其他值失败
		context.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	dao.PublishActionDao(user, videoUrl, picUrl)
	context.JSON(http.StatusOK, response.Response{StatusCode: 0, StatusMsg: picUrl})
}

// 用时间戳混淆时间戳
func fileName(o string) string {
	fileExt := path.Ext(o)
	fName := strings.TrimSuffix(o, fileExt)

	// 毫秒级混淆文件名称
	return fmt.Sprintf("%s-%s%s", fName, time.Now().Format("15:04:05.0000"), fileExt)
}
