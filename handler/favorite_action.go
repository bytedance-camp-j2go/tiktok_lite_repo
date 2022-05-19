package handler

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/service"
	"strconv"
)

func FavoriteAction(userIdStr string, tokenStr string, videoIdStr string, actionTypeStr string) {
	//参数转换
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return
	}
	token, err := strconv.ParseInt(tokenStr, 10, 64)
	if err != nil {
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		return
	}

	//获取service层结果
	service.FavoriteAction(userId, token, videoId, actionType)

}
