package model

import "gorm.io/gorm"

// Video 发布视频
type Video struct {
	gorm.Model
	UserId        int64  `json:"user_id" json:"user_id,omitempty"`               // 上传视频用户id
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`             // 视频播放地址
	CoveUrl       string `json:"cove_url" json:"cove_url,omitempty"`             // 视频封面地址
	FavoriteCount int    `json:"favorite_count" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int    `json:"comment_count" json:"comment_count,omitempty"`   // 视频的评论总数
	Title         string `json:"title" json:"title,omitempty"`                   // 视频标题
}
