package bootstrap

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

// Config >> 初始化配置
// 处理逻辑:
// 1. 读默认配置 `config.yaml` 配置, 如果配置不存在则在当前工作路径创建并输出提示
// 	TODO 如果不存在配置文件, 则从云端配置下载 ( 因为编译运行产出 工作路径可能没有配置文件 / 或者内存放一个默认的配置文件, 进行写入
// 2. 对配置进行重载, 根据 GO_ENV 进行读取,
// 	如果 GO_ENV 不为空, 则执行用 config.{GO_ENV}.yaml 覆盖配置文件, 分环境完成自定义配置读入

var (
	configType  = "yaml"
	defaultPath = "./"
)

func Config() {
	v := viper.New()
	// 从default中读取默认的配置
	v.SetConfigName("config")
	v.AddConfigPath(defaultPath)
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	if err != nil {
		zap.L().Error("err >> ", zap.String("error", err.Error()))
		return
	}

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	initConfig(v)
}

// 尝试读取环境相关联的配置参数
func readEnvConfig(v *viper.Viper) {
	env := os.Getenv("GO_ENV")
	// 根据配置的env读取相应的配置信息
	if env != "" {
		configName := fmt.Sprintf("config.%s", env)
		viper.SetConfigName(configName)
		viper.AddConfigPath(defaultPath)
		viper.SetConfigType(configType)
		err := viper.ReadInConfig()

		// 如果 env 的配置文件不存在, 则创建并提示
		if err != nil {
			dir, _ := os.Getwd()
			profilePath := fmt.Sprintf("%s/%s/%s.%s", dir, defaultPath, configName, configType)
			_, err := util.CreatFile(profilePath)
			if err != nil {
				zap.L().Error(
					"creat env profile error!", zap.Errors("err", []error{err}),
				)
				return
			}
			zap.L().Info(
				"creat env profile : %s\n\tRecommend custom configuration parameters, see 「config.yaml」 to get the configuration item prompt",
				zap.String("filePath", profilePath),
			)
			return
		}
		zap.L().Info(
			"read env config >> %s",
			zap.String("GO_ENV", env),
		)
	}
}

// 将获取到的参数通过序列化转为 config 对象存入全局变量中
func initConfig(v *viper.Viper) {
	serverConfig := config.Config{}
	// 给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	config.Conf = serverConfig
	color.Blue("init success! \n%v", config.Conf)
}
