package response

import (
	"fmt"
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

func (v Video) String() string {
	return fmt.Sprintf("viode {name: %v,ptime: %q,}\n", v.Title, v.CreatedAt.Format("2006-1-2 15:04:05.0000"))

}
