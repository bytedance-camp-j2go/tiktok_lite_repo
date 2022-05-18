package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	bootstrap.Config()
	bootstrap.Logger()
	r := gin.New()
	bootstrap.Router(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		config.Logf.Errorf("serve run error >> %s", err.Error())
		return
	}

	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
