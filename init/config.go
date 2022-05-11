package init

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// Config >> 初始化配置
// 处理逻辑:
// 	1. 读配置, 如果配置不存在则在当前工作路径创建并输出提示
//  2. 会用到三种配置
func Config() {
	// 实例化viper
	v := viper.New()
	// 文件的路径如何设置
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := config.Config{}
	// 给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.Settings = serverConfig
	color.Blue("11111111", global.Settings.LogsAddress)

}
