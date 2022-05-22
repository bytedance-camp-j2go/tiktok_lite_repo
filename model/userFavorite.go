package model

import "gorm.io/gorm"

type UserFavorite struct {
	gorm.Model
	userId  int64 `json:"user_id"`  // 用户id
	videoId int64 `json:"video_id"` // 所点赞的视频id
}
