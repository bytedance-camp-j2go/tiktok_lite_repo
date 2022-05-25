package dao

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

// VideoQueryList 查询视频列表
// 参考 >> db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
func VideoQueryList(videoId []int64) ([]model.Video, error) {
	res := make([]model.Video, len(videoId))
	return res, global.DB.Where("video_id IN ?", videoId).Find(&res).Error
}
