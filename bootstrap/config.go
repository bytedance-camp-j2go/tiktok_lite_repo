package bootstrap

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"os"
	"tiktok-lite/config"
	"tiktok-lite/global"
	"tiktok-lite/util"
)

// InitConfig >> 初始化配置
// 处理逻辑:
// 1. 读默认配置 `config.yaml` 配置, 如果配置不存在则在当前工作路径创建并输出提示
// 2. 对配置进行重载, 根据 GO_ENV 进行读取,
// 	如果 GO_ENV 不为空, 则执行用 config.{GO_ENV}.yaml 覆盖配置文件, 分环境完成自定义配置读入

var (
	configName  = "config"
	configType  = "yaml"
	defaultPath = "./"
)

func InitConfig() {
	v := viper.New()
	// 从default中读取默认的配置
	v.SetConfigName(configName)
	v.AddConfigPath(defaultPath)
	v.SetConfigType(configType)

	// TODO 如果不存在配置自动生成
	if err := v.ReadInConfig(); err != nil {
		panic("read config >> " + err.Error())
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	readEnvConfig()
	init2Config()
}

// 尝试读取环境相关联的配置参数
func readEnvConfig() {
	env := os.Getenv("GO_ENV")
	// 根据配置的env读取相应的配置信息
	if env != "" {
		envConfigName := fmt.Sprintf("config.%s", env)
		viper.SetConfigName(envConfigName)
		viper.AddConfigPath(defaultPath)
		viper.SetConfigType(configType)
		err := viper.ReadInConfig()

		// 如果 env 的配置文件不存在, 则创建并提示
		if err != nil {
			dir, _ := os.Getwd()
			profilePath := fmt.Sprintf("%s/%s/%s.%s", dir, defaultPath, envConfigName, configType)
			_, err := util.CreatFile(profilePath)
			if err != nil {
				color.Red(
					"creat env profile error, The default configuration will be used! >> ", err)
				return
			}
			color.Red(
				"creat env profile : %s"+
					"\n\t"+
					"Recommend custom configuration parameters, "+
					"see 「config.yaml」 to get the configuration item prompt.\n", profilePath)
			return
		}
		fmt.Printf(
			"use env config >> `config.%s.%s`\n",
			env, configType,
		)
	}
}

// 序列化 viper 为 config 对象
func init2Config() {
	serverConfig := config.Config{}
	// 给serverConfig初始值
	if err := viper.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	// 传递给全局变量 对象会被复制
	global.Conf = serverConfig
	color.Blue("init success! \n%v", global.Conf)
}
