package errno

import (
	"fmt"
	"runtime"

	"github.com/petermattis/goid"
)

type Errno struct {
	Code       int64
	Message    string
	file       string
	goid       int64
	InnerErrno *Errno
}

func (e *Errno) Error() string {
	return e.Message
}

func (e *Errno) Unwrap() error {
	return e.InnerErrno
}

func (e *Errno) WithMessage(message string) *Errno {
	e.Message = message
	return e
}

func (e *Errno) WithInnerError(err error) *Errno {
	_, file, line, ok := runtime.Caller(1)
	e.goid = goid.Get()
	if ok {
		e.file = fmt.Sprintf("%s:%d", file, line)
	} else {
		e.file = "unknown"
	}
	e.InnerErrno = ConvertErrno(err)
	return e
}

func (e *Errno) PrintStack() string {
	p := e.InnerErrno
	stack := fmt.Sprintf(
		"Error in goroutine %d, code: %d, message: %s, file: %s\n",
		goid.Get(), e.Code, e.Message, e.file,
	)
	for p != nil {
		if p.Code != 0 {
			stack += fmt.Sprintf(
				"Error in goroutine %d, code: %d, message: %s, file: %s\n",
				p.goid, p.Code, p.Message, p.file,
			)
		} else {
			stack += fmt.Sprintf(
				"Error with message: %s\n",
				p.Message,
			)
		}
		p = p.InnerErrno
	}
	return stack
}

func NewErrno(code int64, message string) *Errno {
	e := &Errno{Code: code, Message: message, InnerErrno: nil}
	_, file, line, ok := runtime.Caller(1)
	goid := goid.Get()
	if ok {
		e.file = fmt.Sprintf("%s:%d", file, line)
	} else {
		e.file = "unknown"
	}
	e.goid = goid
	return e
}

func NewErrnoWithInnerErrno(code int64, message string, innerErrno *Errno) *Errno {
	e := &Errno{Code: code, Message: message, InnerErrno: innerErrno}
	_, file, line, ok := runtime.Caller(1)
	goid := goid.Get()
	if ok {
		e.file = fmt.Sprintf("%s:%d", file, line)
	} else {
		e.file = "unknown"
	}
	e.goid = goid
	return e
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
