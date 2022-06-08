package myerrors

import "fmt"

// -----------------------------------------------------

//MyStringError 简单字符串错误
type MyStringError string

func (e MyStringError) Error() string {
	return fmt.Sprintf("MyStringError %q", string(e))
}

//-----------------------------------------------------
//复杂形错误

//ErrCode 错误代码
type ErrCode uint32

const (
	ErrCodeNo   ErrCode = 0x1
	ErrCodeNull ErrCode = 0x2
)

var errCodeName = map[ErrCode]string{
	ErrCodeNo:   "NO_ERROR",
	ErrCodeNull: "Null_ERROR",
}

func (e ErrCode) String() string {
	if s, ok := errCodeName[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown myerrors code 0x%x", uint32(e))
}

func (e ErrCode) stringToken() string {
	if s, ok := errCodeName[e]; ok {
		return s
	}
	return fmt.Sprintf("ERR_UNKNOWN_%d", uint32(e))
}

//ConnectionError 使用变量做错误（good~~！！）
type ConnectionError ErrCode

//Error 自定义变量的方法。变量也可以有自己的方法。
func (e ConnectionError) Error() string {
	return fmt.Sprintf("Connection Error: %s", ErrCode(e))
}

//-----------------------------------------------------
