package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/init"
	"github.com/gin-gonic/gin"
)

func init() {
	// 这里进行初始化工作...

	init.Config()
	// 初始化完成以后，global.XXX 可以被使用
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		// TODO 由统一日志输出 错误信息
		return
	}
}
