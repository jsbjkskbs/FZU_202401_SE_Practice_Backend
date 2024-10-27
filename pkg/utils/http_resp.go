package utils

import (
	"sfw/pkg/errno"
)

// BaseHttpResponse 基础Http响应结构
type BaseHttpResponse struct {
	Code int64
	Msg  string
}

func baseHttpResponse(err errno.Errno) *BaseHttpResponse {
	return &BaseHttpResponse{
		Code: err.Code,
		Msg:  err.Message,
	}
}

// CreateBaseHttpResponse 创建基础Http响应
func CreateBaseHttpResponse(err error) *BaseHttpResponse {
	if err == nil {
		return baseHttpResponse(*errno.NoError)
	}

	if e, ok := err.(*errno.Errno); ok {
		return baseHttpResponse(*e)
	}

	s := errno.CustomError.WithMessage(err.Error())
	return baseHttpResponse(*s)
}
