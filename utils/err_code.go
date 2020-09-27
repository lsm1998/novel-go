package utils

const (
	OK = 200 + iota
	// 请求参数错误
	ERR_REQUEST_PARAMS
	// 用户不存在
	ERR_USER_NOT_EXIS
	// 用户禁用
	ERR_USER_DISABLE
	// 限流
	ERR_LIMIT
	// token校验失败
	ERR_TOKEN_VALID
	// 登录失败，账号或者密码错误
	ERR_LOGIN_FAIL
	// 登录失败，重复登录
	ERR_LOGIN_REPEAT
	// 长连接失败
	ERR_CONN
	// 数据库错误
	ERR_DB
	// 服务端错误
	ERR_SERVER_ERR
)

type Error struct {
	Code int
	Msg  string
}

func (e Error) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{code, msg}
}
