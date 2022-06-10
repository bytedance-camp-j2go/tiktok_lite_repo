package dao

import (
	"tiktok-lite/global"
	"tiktok-lite/model"
)

// VideoQueryList 查询视频列表
// 参考 >> db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
func VideoQueryList(videoId []int64) ([]model.Video, error) {
	res := make([]model.Video, len(videoId))
	if err := global.DB.Where("video_id IN ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}
	videoProcess(res)
	return res, nil
}

// video process 一些数据字段存在缓存、其他表中的，将其取出
func videoProcess(videos []model.Video) {
	for idx := range videos {
		videos[idx].FavoriteCount = GetFavoriteCountByVideoId(videos[idx].VideoId)
		videos[idx].CommentCount = CommentCnt(videos[idx].VideoId)

	}
}
