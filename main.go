package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//加载配置
	bootstrap.InitConfig()
	//加载日志
	bootstrap.InitLogger()
	//mysql初始化，初始化连接对象
	bootstrap.InitDB()
	//获取路由
	r := gin.New()
	r.Use(middleware.JWTAuth())
	//初始化路由
	bootstrap.InitRouter(r)
	//初始化数据库
	bootstrap.InitDB()
	err := r.Run(fmt.Sprintf(":%d", global.Conf.Port))
	if err != nil {
		global.Logf.Errorf("serve run error >> %s", err.Error())
		return
	}
	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
