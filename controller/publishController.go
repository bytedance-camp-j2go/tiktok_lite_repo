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
	"github.com/wxnacy/wgo/arrays"
	"go.uber.org/zap"
	"net/http"
	"path"
	"strings"
	"time"
)

// PublishAction 视频投稿
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
	// 获取视频标题
	title := context.Query("title")
	if err != nil {
		// 注意：状态码 0成功 其他值失败
		context.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	dao.PublishActionDao(user, videoUrl, picUrl, title)
	context.JSON(http.StatusOK, response.Response{StatusCode: 0, StatusMsg: picUrl})
}

// 用时间戳混淆时间戳
func fileName(o string) string {
	fileExt := path.Ext(o)
	fName := strings.TrimSuffix(o, fileExt)

	// 毫秒级混淆文件名称
	return fmt.Sprintf("%s-%s%s", fName, time.Now().Format("15:04:05.0000"), fileExt)
}

// PublishList 发布列表
// 场景：登录用户的视频发布列表，列出用户所有投稿过的视频
func PublishList(context *gin.Context) {
	// 获取用户信息
	var a any
	a, _ = context.Get(global.UserName)
	user := a.(model.User)
	// 封装用户响应信息
	userResp := response.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true, // 注意；这里是用户看自己的主页，所以肯定是关注了自己的
	}
	// 查询用户发布的视频列表
	videos, err := dao.PublishList(user.Id)
	if err != nil {
		context.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "查询失败"})
		return
	}
	// 查询用户点赞过自己的视频
	videosId, _ := dao.UserFavorite(user.Id)
	size := len(videos)
	videosResp := make([]response.VideoList, size, size)
	// var videosResp [size]response.VideoList
	// 创建响应对象
	for i, v := range videos {
		videosResp[i].Id = int64(v.Model.ID)
		videosResp[i].Author = userResp // 用户信息
		videosResp[i].PlayUrl = v.PlayUrl
		videosResp[i].CoveUrl = v.CoverUrl
		videosResp[i].FavoriteCount = v.FavoriteCount
		videosResp[i].CommentCount = v.CommentCount
		videosResp[i].CommentCount = v.CommentCount
		videosResp[i].IsFavorite = arrays.ContainsInt(videosId, int64(v.ID)) > 0
		videosResp[i].IsFavorite = true // 注意：这块需要判断用户对这个视频有没有点赞
		videosResp[i].Title = v.Title
	}
	context.JSON(http.StatusOK, response.PublishListResponse{
		Response:  response.Response{StatusCode: 0, StatusMsg: "成功"},
		VideoList: videosResp,
	})
}
