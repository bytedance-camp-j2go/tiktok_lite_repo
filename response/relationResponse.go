package response

type RelationActionResponse struct {
	Response
}

type RelationFollowListResponse struct {
	Response
	UserList []RelationUser `json:"user_list"`
}

type RelationFollowerListResponse struct {
	Response
	UserList []RelationUser `json:"user_list"`
}

type RelationUser struct {
	Id            int64  `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户昵称
	FollowCount   int64  `json:"follow_count"`   // 用户关注人数
	FollowerCount int64  `json:"follower_count"` // 用户粉丝数量
	IsFollow      bool   `json:"is_follow"`      //对方是否关注
}
