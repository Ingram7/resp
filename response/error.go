package response

import "encoding/json"

var _ Error = (*err)(nil)

type Error interface {
	error
	// WithErr 设置错误信息
	WithErr(err error) Error
	// GetCode 获取 Business Code
	GetCode() int
	// GetMsg 获取 Msg
	GetMsg() string
	// GetErr 获取错误信息
	GetErr() error
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	Code int    // Business Code
	Msg  string // 描述信息
	Err  error  // 错误信息
}

func NewErrorWithMsg(Code int, msg string) Error {
	return &err{
		Code: Code,
		Msg:  msg,
	}
}

func NewError(Code int) Error {
	return &err{
		Code: Code,
		Msg:  Text(Code),
	}
}

func (e *err) Error() string {
	return e.ToString()
}

// WithErr 封装真实 err
func (e *err) WithErr(err error) Error {
	e.Err = err
	return e
}

// GetCode 获取 Code
func (e *err) GetCode() int {
	return e.Code
}

// GetMsg 获取 Msg
func (e *err) GetMsg() string {
	return e.Msg
}

// GetErr 获取 Err
func (e *err) GetErr() error {
	return e.Err
}

func (e *err) ToString() string {
	err := &struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{
		Code: e.Code,
		Msg:  e.Msg,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
