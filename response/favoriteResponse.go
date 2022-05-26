package response

import "github.com/bytedance-camp-j2go/tiktok_lite_repo/model"

// FavoriteActionResponse 用户点赞返回响应
type FavoriteActionResponse struct {
	Response
}

// FavoriteListResponse 用户获取点赞列表响应
type FavoriteListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}
