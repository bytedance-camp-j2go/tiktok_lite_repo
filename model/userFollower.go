package model

type UserFollower struct {
	Id             int64 //表主键id
	UserId         int64 `gorm:"column:user_id"`          //用户id
	FollowerUserId int64 `gorm:"column:follower_user_id"` //关注的粉丝id
}
