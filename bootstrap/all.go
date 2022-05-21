package bootstrap

import (
	"sync"
)

// InitAll 执行项目整体的初始化
// 除了 router 路由信息都可以用 InitAll() 初始化
//
/*	主体逻辑
	// 加载配置
	bootstrap.InitConfig()
	// 加载日志
	bootstrap.InitLogger()
	// mysql初始化，初始化连接对象
	bootstrap.InitDB()
	// 自动迁移数据库表
	bootstrap.InitModel()
	bootstrap.InitRedis()
*/
func InitAll() {
	InitConfig()
	InitLogger()
	// 思考: redis 和 db 可以并发执行初始化, 想要简化 init 写法
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		InitRedis()
		wg.Done()
	}()

	// 初始化数据库, 迁移模型
	go func() {
		InitDB()
		go InitModel()
		wg.Done()
	}()
	wg.Wait()

	// 初始化存储器
	InitDriver()
}
