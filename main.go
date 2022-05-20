package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"os"
)

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
	bootstrap.InitConfig()
	bootstrap.InitLogger()
	r := gin.New()
	bootstrap.InitRouter(r)

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

	err = r.Run(fmt.Sprintf(":%d", global.Conf.Port))
	if err != nil {
		global.Logf.Errorf("serve run error >> %s", err.Error())
		return
	}

	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
