package response

// PublishActionResponse 视频投稿响应
type PublishActionResponse struct {
	Response
}

// VideoList 用户发表列表中视频列表结构
type VideoList struct {
	Id            int64  `json:"id"`             // 视频唯一标识
	Author        User   `json:"author"`         // 视频用户信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoveUrl       string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int    `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int    `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // 是否点赞
	Title         string `json:"title"`          // 视频标题
}

// PublishListResponse 发布列表响应
type PublishListResponse struct {
	Response              // 状态码、状态信息
	VideoList []VideoList `json:"video_list"` // 视频列表
}
