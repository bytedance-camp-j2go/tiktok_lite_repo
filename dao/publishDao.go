package dao

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"gorm.io/gorm"
	"time"
)

// PublishActionDao 视频投稿，将视频信息持久化到数据库中
func PublishActionDao(user model.User, playUrl string, coverUrl string, title string) (int64, error) {
	db := global.DB
	video := model.Video{
		VideoId:       util.UniqueID(),
		Model:         gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		UserId:        user.Id,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	err := db.Create(video).Error

	if err != nil {
		return 0, err
	}

	return video.VideoId, nil
}

// PublishList 查询用户发布视频列表
func PublishList(userId int64) ([]model.Video, error) {
	db := global.DB
	var videos []model.Video
	err := db.Where("user_id=?", userId).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

// UserFavorite 用户点赞的视频列表
func UserFavorite(userId int64) ([]int64, error) {
	db := global.DB
	var videosId []int64
	err := db.Table("user_favorite").Select("video_id").Where("user_id=?", userId).Find(&videosId).Error
	if err != nil {
		return videosId, err
	}
	return videosId, nil

}
