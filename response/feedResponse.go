package response

// 视频流接口的响应

type FeedResponse struct {
	StatusCode int `json:"status_code"`
	VideoList  []struct {
		Id       int    `json:"id"`
		User     User   `json:"author"`
		PlayUrl  string `json:"play_url"`
		CoverUrl string `json:"cover_url"`
	} `json:"video_list"`
	NextTime int `json:"next_time"`
}
