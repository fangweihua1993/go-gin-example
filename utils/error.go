// @Description
// @Author yixia
// @Copyright 2021 sndks.com. All rights reserved.
// @LastModify 2021/1/14 5:21 下午

package utils

import (
	"errors"
	"fmt"
)

var (
	// ErrNoDefined 未定义
	ErrNoDefined = errors.New("error no defined")
	// ErrNotImplemented 未实现
	ErrNotImplemented = errors.New("not implemented")
	// ErrNotExist 记录不存在
	ErrNotExist = errors.New("record does not exist")
	// ErrZeroRowsAffected zero row affect
	ErrZeroRowsAffected = errors.New("zero rows affect")
)

type Error struct {
	Code    int
	Msg     string
	nestErr error
}

// NewCustomError
//  @Description 创建自定义错误
//  @param code 错误码
//  @param msg 业务错误消息，供业务提示使用
//  @param ess 嵌套错误切片
//  @return error 返回 Error 类型的错误
func NewCustomError(code int, msg string, ess ...error) error {
	return doNewCustomError(code, msg, ess...)
}

func doNewCustomError(code int, msg string, ess ...error) Error {
	if len(ess) == 0 {
		return Error{Code: code, Msg: msg}
	}
	var wrapErr error
	for i := range ess {
		wrapErr = fmt.Errorf("%w", ess[i])
	}
	customErr := Error{Code: code, Msg: msg}.WithError(wrapErr)
	return customErr
}

// NewError
// 	@Description
//	@Param code
//	@Param err
// 	@Return error
func NewError(code int, ess ...error) error {
	//return Error{Code: code}
	return doNewCustomError(code, "", ess...)
}

// WrapError
//  @Description 包装 错误
//  @param e Error 类型错误，
//  @param nestError 嵌套错误
//  @return error 如e 是 Error 类型，修改 nestErr 为 nestError,否则返回 e
func WrapError(e error, nestError error) error {
	if v, ok := e.(Error); ok {
		v.nestErr = nestError
		return v
	}
	return e
}

// GetError
//  @Description 获取原始 Error
//  @param e 原始 error
//  @return *Error 如 Error 就返回
func GetError(e error) *Error {
	if v, ok := e.(Error); ok {
		return &v
	}
	if e == nil {
		return nil
	}
	return &Error{Msg: e.Error()}
}

// Error
// 	@Description
// 	@Receiver e
// 	@Return string
func (e Error) Error() string {
	return e.Msg
}

// WithError
//  @Description 设置 error
//  @receiver e
//  @param err
//  @return error
func (e Error) WithError(err error) Error {
	e.nestErr = err
	return e
}

// GetNestError
//  @Description 获取嵌套错误
//  @receiver e
//  @return error
func (e Error) GetNestError() error {
	return e.nestErr
}

// Unwrap
// 	@Description 实现 wrapper ，被 errors.Unwrap() 调用
// 	@Receiver e
// 	@Return error 如果没有实现 wrapper，返回 e 中的嵌套错误
func (e Error) Unwrap() error {
	u, ok := e.nestErr.(interface {
		Unwrap() error
	})
	if !ok {
		return e.nestErr
	}
	return u.Unwrap()
}
