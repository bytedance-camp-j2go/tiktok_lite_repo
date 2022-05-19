package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	//加载配置
	bootstrap.Config()
	//加载日志
	bootstrap.Logger()
	//mysql初始化，初始化连接对象
	bootstrap.InitDB()
	//获取路由
	r := gin.New()
	//初始化路由
	bootstrap.Router(r)

	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		config.Logf.Errorf("serve run error >> %s", err.Error())
		return
	}

	fmt.Print("main >> ")
	fmt.Println(os.Getwd())
}
