package errno

import (
	"fmt"
)

type Errno struct {
	Code       int64
	Message    string
	InnerErrno *Errno
}

func (e *Errno) Error() string {
	return fmt.Sprintf("Error Message: %s", e.Message)
}

func (e *Errno) Unwrap() error {
	return e.InnerErrno
}

func (e *Errno) WithMessage(message string) *Errno {
	e.Message = message
	return e
}

func (e *Errno) PrintStack() string {
	p := e.InnerErrno
	stack := fmt.Sprint("main error: ", e.Error(), "\n")
	for p != nil {
		stack += "\t" + e.Error() + "\n"
		p = p.InnerErrno
	}
	return stack
}

func NewErrno(code int64, message string) *Errno {
	return &Errno{Code: code, Message: message, InnerErrno: nil}
}

func NewErrnoWithInnerErrno(code int64, message string, innerErrno *Errno) *Errno {
	return &Errno{Code: code, Message: message, InnerErrno: innerErrno}
}

func ConvertErrno(err error) *Errno {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Errno); ok {
		return e
	}
	return NewErrno(0, err.Error())
}
