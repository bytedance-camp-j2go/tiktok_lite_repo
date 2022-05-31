package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"tiktok-lite/bootstrap"
	"tiktok-lite/global"
	"tiktok-lite/middleware"
)

func main() {
	bootstrap.InitAll()
	// 获取路由
	r := gin.New()
	r.Use(middleware.JWTAuth())
	// 初始化路由
	bootstrap.InitRouter(r)

	err := r.Run(fmt.Sprintf(":%d", global.Conf.Port))
	if err != nil {
		global.Logf.Errorf("serve run error >> %s\n", err.Error())
		return
	}
	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
