package response

// User 返回用户信息中的User对象
type User struct {
	Id            int64  `json:"id"`             // 视频发布者id，这个id就是user_id
	Name          string `json:"name"`           // 视频发布者昵称
	FollowCount   int64  `json:"follow_count"`   // 视频发布者关注数量
	FollowerCount int64  `json:"follower_count"` // 视频发布者粉丝数量
	IsFollow      bool   `json:"is_follow"`      // 用户是否关注这个视频发布者
}

// UserResponse 返回用户信息
type UserResponse struct {
	Response
	User User `json:"user"`
}

// UserLoginResponse 用户登录成功返回响应
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 鉴权token
}

// UserRegisterResponse 用户注册成功返回响应
type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 鉴权token
}
