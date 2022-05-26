package global

// 定义日志等级, 和 zap 日志等级一一对应
// 啰嗦一句: fatal 和 panic 区别, fatal 会让程序直接终止 (见 fmt.Fatal 源码终端 os.Exit(1)
// panic 更类似于 Java 之 Exception, 会逐级上抛, 通过 defer recover 接受 panic 并处理, 基本等同于 try-catch
const (
/*LevelDebug = iota
  LevelInfo
  LevelWaring
  LevelError
  LevelDPanic
  LevelPanic
  LevelFatal*/
)

const (
/*// DefExpiration
  DefExpiration = 0*/
)

const (
	// CtxUserKey 鉴权成功后，在context中存入user信息，其key为userName
	CtxUserKey = "requestUser"

	// VideoSeqSetKey 视频列表 ZSet key
	VideoSeqSetKey = "video:id-seq"
)
