package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	bootstrap.InitConfig()
	bootstrap.InitLogger()
	r := gin.New()
	bootstrap.InitRouter(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(fmt.Sprintf(":%d", global.Conf.Port))
	if err != nil {
		global.Logf.Errorf("serve run error >> %s", err.Error())
		return
	}

	fmt.Print("main > ")
	fmt.Println(os.Getwd())
}
