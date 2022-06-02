package response

// 用户点赞返回响应
type FavoriteActionResponse struct {
	Response
}

// 用户获取点赞列表响应
type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// type FavoriteVideo struct {
// 	VideoId       int64        `json:"id"`
// 	User          FavoriteUser `json:"author"`         // 上传视频用户id
// 	PlayUrl       string       `json:"play_url"`       // 视频播放地址
// 	CoverUrl      string       `json:"cover_url"`      // 视频封面地址
// 	FavoriteCount int64        `json:"favorite_count"` // 视频的点赞总数
// 	CommentCount  int          `json:"comment_count"`  // 视频的评论总数
// 	IsFavorite    bool         `json:"is_favorite"`    // 视频是否点赞
// 	Title         string       `json:"title"`          // 视频标题
// }
//
// type FavoriteUser struct {
// 	Id            int64  `json:"id"`             // 用户id
// 	Name          string `json:"name"`           // 用户昵称
// 	FollowCount   int64  `json:"follow_count"`   // 用户关注人数
// 	FollowerCount int64  `json:"follower_count"` // 用户粉丝数量
// 	IsFollow      bool   `json:"is_follow"`      // 对方是否关注
//
// }
