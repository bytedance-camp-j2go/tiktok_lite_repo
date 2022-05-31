package model

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:id" gorm:"primary_key"`  // 用户id
	UserName      string `json:"user_name" gorm:"column:username"`                  // 用户登录帐号
	PassWord      string `json:"password,omitempty" gorm:"column:password"`         // 用户密码
	Name          string `json:"name,omitempty" gorm:"column：name"`                 // 用户昵称
	FollowCount   int64  `json:"follow_Count,omitempty" gorm:"column:follow_count"` // 用户关注人数
	FollowerCount int64  `json:"follower_count" gorm:"column:follower_count"`       // 用户粉丝数量

	// 随机图 API 用作背景 URL
	// https://api.lixingyong.com/api/images
	// # 随机头像API请求方式 #
	// Method: GET
	// # 请求地址 #
	// https://api.sunweihu.com/api/sjtx/api.php
	// # 参数 #
	// lx【1.男头：a1 2.女头：b1 3.动漫：c1 4.动漫女头：c2 5.动漫男头：c3】
	// # 返回数据 #
	// 本API无返回数据,直接输出头像
	// # 备注 #
	// 后续将会增加更多类型的头像，欢迎大家投稿
	// # 示例 #
	// https://api.sunweihu.com/api/sjtx/api.php?lx=c1

}
