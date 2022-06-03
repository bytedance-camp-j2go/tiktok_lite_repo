package dao

import (
	"go.uber.org/zap"
	"tiktok-lite/global"
	"tiktok-lite/model"
)

// var (
// 	BadComment = &model.Comment{}
// )

// CommentQuery 读单条记录，如果是
func CommentQuery(id int64) (*model.Comment, error) {
	res := model.Comment{}
	if err := global.DB.Where("id = ?", id).First(&res).Error; err != nil {
		zap.L().Debug("query comment error", zap.Error(err))
		return nil, nil
	}

	return &res, nil
}

func CommentQueryUserId(id int64) (int64, error) {
	var res int64
	err := global.DB.
		Table("comment").
		Select("user_id").
		Where("id = ?", id).
		First(&res).Error
	return res, err
}

func CommentQueryList(ids []int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := global.DB.Where("id IN ?", ids).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func CommentDel(comment *model.Comment) error {
	return global.DB.Delete(comment).Error
}

func CommentSave(comment *model.Comment) error {
	return global.DB.Save(comment).Error
}

func CommentList(videoId int64) ([]model.Comment, error) {
	var commentList []model.Comment
	err := global.DB.Where("video_id = ?", videoId).Find(&commentList).Error
	return commentList, err
}
