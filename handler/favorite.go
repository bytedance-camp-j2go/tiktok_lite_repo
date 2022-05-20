package handler

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func FavoriteAction(c *gin.Context) {
	userId := c.Query("user_id")
	//token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	//todo 鉴定token是否有效
	actionTypeInt, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		fmt.Println("行为类型无效")
		return
	}

	//行为类型为1，点赞
	//行为类型为2，取消点赞
	if actionTypeInt == 1 {
		zset := &redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: videoId,
		}
		err = global.RedisDB.ZAdd(c, userId, zset).Err()

		if err != nil {
			fmt.Printf("redis sadd failed! err:%v\n", err)
			return
		}
		fmt.Println("点赞成功")
	} else if actionTypeInt == 2 {
		err = global.RedisDB.ZRem(c, userId, videoId).Err()
		if err != nil {
			fmt.Printf("redis zrem failed! err:%v\n", err)
			return
		}
		fmt.Println("取消点赞成功")
	} else {
		fmt.Println("未知的行为类型")
	}

}

func FavoriteList(c *gin.Context) {
	userId := c.Query("user_id")
	//token := c.Query("token")
	//todo 鉴定token是否有效

	res, err := global.RedisDB.ZRange(c, userId, 0, -1).Result()
	if err != nil {
		fmt.Printf("redis smembers failed! err:%v\n", err)
		return
	}
	fmt.Println(res)

	//todo 通过视频id数组 查询到 视频对象数组
}
