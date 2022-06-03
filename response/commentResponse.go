package response

import (
	"go.uber.org/zap"
	"tiktok-lite/model"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type Comment struct {
	model.Comment
	CreatedDate string `json:"created_date,omitempty"`
	User        User   `json:"user,omitempty"`
}

const (
	commentDateFormat = "01-02"
)

func NewComment(comment model.Comment, u2id int64) Comment {
	res := Comment{
		Comment:     comment,
		CreatedDate: comment.CreatedAt.Format(commentDateFormat),
	}

	if user, err := NewUser(comment.UsrID, u2id); err == nil {
		res.User = user
	} else {
		go zap.L().Debug("query comment error!",
			zap.Int64("cid", comment.ID),
			zap.Int64("u2id", u2id),
			zap.Error(err),
		)
	}

	return res
}

func NewCommentListResp(cs []model.Comment, u2id int64) CommentListResponse {
	res := make([]Comment, 0, len(cs))
	for _, c := range cs {
		res = append(res, NewComment(c, u2id))
	}

	return CommentListResponse{
		Response:    BaseSuccess(""),
		CommentList: res,
	}
}

func NewCommentResp(c model.Comment, u2id int64) CommentResponse {
	return CommentResponse{
		Response: BaseSuccess("comment success"),
		Comment:  NewComment(c, u2id),
	}
}
