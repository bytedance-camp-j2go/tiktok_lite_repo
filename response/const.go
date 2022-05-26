package response

// StatusCode 状态枚举
// 除了 0 以外的 status code 都表示操作出错
const (
	OptSuccess = iota
	OptInputError
	_
	_
	_
	OptServerError
)

var (
	// UserError 通过两个 id 计算如果错误返回这个
	UserError = User{}
)
