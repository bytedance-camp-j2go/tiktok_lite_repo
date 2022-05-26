package model

import "gorm.io/gorm"

type UserFavorite struct {
	gorm.Model
	UserId  int64 `json:"user_id" gorm:"index"` // 用户id
	VideoId int64 `json:"video_id"`             // 所点赞的视频id
}
