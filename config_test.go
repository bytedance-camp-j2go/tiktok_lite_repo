package main_test

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/bootstrap"
	"go.uber.org/zap"
	"os"
	"testing"
)

// 多版本测试
func TestMultiVersionViper(t *testing.T) {
	bootstrap.InitConfig()
	bootstrap.InitLogger()
	zap.L().Debug("111")
}

// 测试环境变量 `GO_ENV` 的获取
func TestGetEnv(t *testing.T) {
	env := os.Getenv("GO_ENV")
	fmt.Printf("GO_ENV >> %#q\n", env)
	if env == "" {
		fmt.Println("def")
	}
}
