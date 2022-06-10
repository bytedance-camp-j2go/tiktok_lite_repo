package response

// FavoriteActionResponse 用户点赞返回响应
type FavoriteActionResponse struct {
	Response
}

// FavoriteListResponse 用户获取点赞列表响应
type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
