package model

type User struct {
	Id int64 `json:"id,omitempty" gorm:"column:id" gorm:"primary_key"` // 主键id
	// UserId        int64  `json:"user_id" gorm:"column:user_id"`                     // 用户id
	UserName      string `json:"user_name" gorm:"column:username"`                  // 用户登录帐号
	PassWord      string `json:"password,omitempty" gorm:"column:password"`         // 用户密码
	Name          string `json:"name,omitempty" gorm:"column：name"`                 // 用户昵称
	FollowCount   int64  `json:"follow_Count,omitempty" gorm:"column:follow_count"` // 用户关注人数
	FollowerCount int64  `json:"follower_count" gorm:"column:follower_count"`       // 用户粉丝数量
}
