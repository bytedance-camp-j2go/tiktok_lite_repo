package config

// 定义日志等级, 和 zap 日志等级一一对应
// 啰嗦一句: fatal 和 panic 区别, fatal 会让程序直接终止 (见 fmt.Fatal 源码终端 os.Exit(1)
// panic 更类似于 Java 之 Exception, 会逐级上抛, 通过 defer recover 接受 panic 并处理, 基本等同于 try-catch
const (
	LevelDebug = iota
	LevelInfo
	LevelWaring
	LevelError
	LevelDPanic
	LevelPanic
	LevelFatal
)
