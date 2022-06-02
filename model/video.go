package model

import "gorm.io/gorm"

// Video 发布视频
type Video struct {
	gorm.Model `json:"-"`
	VideoId    int64 `json:"id" gorm:"uniqueIndex:video_id_idx"`
	UserId     int64 `json:"user_id"` // 上传视频用户id

	// 以下计数可以考虑放到 redis 中实现
	FavoriteCount int64 `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64 `json:"comment_count"`  // 视频的评论总数

	PlayUrl  string `json:"play_url"`  // 视频播放地址
	CoverUrl string `json:"cover_url"` // 视频封面地址
	Title    string `json:"title"`     // 视频标题
}
