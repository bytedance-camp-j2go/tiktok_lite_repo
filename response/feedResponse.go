package response

import (
	"tiktok-lite/model"
)

// FeedResponse 视频流接口的响应
type FeedResponse struct {
	StatusCode int     `json:"status_code"`
	VideoList  []Video `json:"video_list"`
	NextTime   int64   `json:"next_time"`
}

type Video struct {
	model.Video
	Author     User `json:"author,omitempty"`
	IsFavorite bool `json:"is_favorite"`
}
