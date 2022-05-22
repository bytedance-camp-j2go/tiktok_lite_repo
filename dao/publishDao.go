package dao

import (
	"container/list"
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"gorm.io/gorm"
	"time"
)

// PublishActionDao 视频投稿，将视频信息持久化到数据库中
func PublishActionDao(user model.User, playUrl string, coverUrl string, title string) error {
	db := global.DB
	video := model.Video{
		Model:         gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		UserId:        user.Id,
		PlayUrl:       playUrl,
		CoveUrl:       coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	err := db.Create(video).Error
	if err != nil {
		return err
	}
	return nil
}

// PublishList 查询用户发布视频列表
func PublishList(userId int64) list.List {
	db := global.DB
	video := model.Video{}
	fmt.Println(video, db)
	return list.List{}
}
