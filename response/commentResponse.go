package response

import (
	"tiktok-lite/model"
)

type CommentResponse struct {
	model.Comment
	User User `json:"user,omitempty"`
}
