package model

import "gorm.io/gorm"

// Video 发布视频
type Video struct {
	gorm.Model
	UserId        int64  `json:"user_id"`        //上传视频用户id
	PlayUrl       string `json:"play_url"`       //视频播放地址
	coveUrl       string `json:"cove_url"`       //视频封面地址
	FavoriteCount int    `json:"favorite_count"` //视频的点赞总数
	CommentCount  int    `json:"comment_count"`  //视频的评论总数
	Title         string `json:"title"`          //视频标题
}
