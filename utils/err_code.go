package utils

const (
	OK = 0
	// 用户不存在
	ERR_NOT_EXIS = 200 + iota
	// token校验失败
	ERR_TOKEN_VALID
	// 登录失败，账号或者密码错误
	ERR_LOGIN_PASS
	// 登录失败，重复登录
	ERR_LOGIN_REPEAT
	// 长连接失败
	ERR_CONN
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
