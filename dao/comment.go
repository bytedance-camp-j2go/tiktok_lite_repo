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

// func CommentQueryList(ids []int64) ([]*model.Comment, error) {
//
// }

// func CommentDel(comment *model.Comment) error {
// 	global.DB.Model(comment).
// }
