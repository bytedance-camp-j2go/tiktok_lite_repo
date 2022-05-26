package response

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

type RelationActionResponse struct {
	Response
}

type RelationFollowListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}

type RelationFollowerListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}
