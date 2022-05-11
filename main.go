package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func init() {
	// 这里进行初始化工作...

	bootstrap.Config()
	bootstrap.Logger()
	// 初始化完成以后，global.XXX 可以被使用
}

func main() {
	r := gin.New()
	bootstrap.Router(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		zap.L().Error("serve run error", zap.String("error", err.Error()))
		return
	}

	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
