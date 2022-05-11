package main

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"testing"
)

// 多版本测试
func TestMultiVersionViper(t *testing.T) {
	bootstrap.Logger()

	configType := "yaml"
	defaultPath := "./"
	v := viper.New()
	// 从default中读取默认的配置
	v.SetConfigName("config")
	v.AddConfigPath(defaultPath)
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("err >> ", err, "\nconfig >>", v)
		return
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	env := os.Getenv("GO_ENV")
	// 根据配置的env读取相应的配置信息
	if env != "" {
		configName := fmt.Sprintf("config.%s", env)
		viper.SetConfigName(configName)
		viper.AddConfigPath(defaultPath)
		viper.SetConfigType(configType)
		err = viper.ReadInConfig()

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
	fmt.Println("err >> ", err, "\nconfig >>", v)
}

// 测试环境变量 `GO_ENV` 的获取
func TestGetEnv(t *testing.T) {
	env := os.Getenv("GO_ENV")
	fmt.Println("GO_ENV >>", env)
	if env != "" {
		fmt.Println(env)
	} else {
		fmt.Println("def")
	}
}
