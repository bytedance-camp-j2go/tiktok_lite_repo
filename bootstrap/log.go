package bootstrap

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/config"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"go.uber.org/zap"
)

// Logger 初始化 log 日志, 使用 zap 的原因是因为高性能
// learn from [juejin.cn](https://juejin.cn/post/6971217119379718175)
func Logger() {
	// 实例化zap 配置
	cfg := zap.NewDevelopmentConfig()

	// 文件日志输出目录在 global.Conf 中配置
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog-%s.log", config.Conf.LogsAddress, util.GetNowFormatTodayTime()), //
		"stdout",
	}

	// 创建logger实例
	logger, _ := cfg.Build()

	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	config.Log = logger        // 注册到全局变量中
}
