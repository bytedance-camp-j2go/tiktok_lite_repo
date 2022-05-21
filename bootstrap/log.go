package bootstrap

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/util"
	"go.uber.org/zap"
)

// InitLogger 初始化 log 日志, 使用 zap 的原因是因为高性能
// learn from [juejin.cn](https://juejin.cn/post/6971217119379718175)
func InitLogger() {
	// 实例化zap 配置
	cfg := zap.NewDevelopmentConfig()

	// level := zap.NewAtomicLevel()
	// initLogLevel(&level, global.Conf.LogLevel)

	// 文件日志输出目录在 config.Conf 中配置
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("%s/%s.log", global.Conf.LogsAddress, util.GetNowFormatTodayTime()), //
		"stdout",
	}

	if _, err := util.PathExists(global.Conf.LogsAddress); err != nil {
		panic(err)
	}

	// 创建logger实例
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)   // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	global.Log = logger          // 注册到全局变量中
	global.Logf = logger.Sugar() // 注册到全局变量中
	global.Log.Debug("logger init success!")
}

// func initLogLevel(atomicLevel *zap.AtomicLevel, logLevel int) {
// 	switch logLevel {
// 	case global.LevelDebug:
// 		atomicLevel.SetLevel(zapcore.DebugLevel)
// 	case global.LevelInfo:
// 		atomicLevel.SetLevel(zapcore.InfoLevel)
// 	case global.LevelWaring:
// 		atomicLevel.SetLevel(zapcore.WarnLevel)
// 	case global.LevelError:
// 		atomicLevel.SetLevel(zapcore.ErrorLevel)
// 	case global.LevelDPanic:
// 		atomicLevel.SetLevel(zapcore.DPanicLevel)
// 	case global.LevelPanic:
// 		atomicLevel.SetLevel(zapcore.PanicLevel)
// 	case global.LevelFatal:
// 		atomicLevel.SetLevel(zapcore.FatalLevel)
// 	}
// }
