package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
)

func init() {
	// 这里进行初始化工作...

	bootstrap.Config()
	bootstrap.Logger()
	// 初始化完成以后，global.XXX 可以被使用
}

func initClient() (err error) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := gin.New()
	bootstrap.Router(r)

	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
	fmt.Println("Redis连接成功")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.PUT("/douying/favorite/action", func(c *gin.Context) {
		userId, _ := c.GetPostForm("user_id")
		token, _ := c.GetPostForm("token")
		videoId, _ := c.GetPostForm("video_id")
		actionType, _ := c.GetPostForm("action_type")
		fmt.Println(userId, token, videoId, actionType)
		handler.FavoriteAction(userId, token, videoId, actionType)
		c.JSON(200, "点赞成功")
	})

	err = r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		zap.L().Error("serve run error", zap.String("error", err.Error()))
		return
	}

	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
