package errno

import (
	"errors"
	"testing"
)

// TestErrno1 test errno
// unsafe operation
func TestErrno1(t *testing.T) {
	var err error
	err = NewErrnoWithInnerErrno(1, "test", InternalServerError)
	t.Log(err.(*Errno).PrintStack())
}

// TestErrno2 test errno
// safe operation
func TestErrno2(t *testing.T) {
	var err error
	err = errors.New("test")
	t.Log(ConvertErrno(err).PrintStack())
}

// TestErrno3 test errno
// safe operation
func TestErrno3(t *testing.T) {
	var err error
	err = NewErrnoWithInnerErrno(1, "test", InternalServerError)
	t.Log(ConvertErrno(err).PrintStack())
}

func TestErrno4(t *testing.T) {
	var err error
	err = NewErrno(1, "test")
	err = err.(*Errno).WithInnerError(errors.New("test inner error"))
	t.Log(err.(*Errno).PrintStack())
}

func TestErrno5(t *testing.T) {
	var err error
	err = errors.New("test")
	err = ConvertErrno(err).WithInnerError(errors.New("test inner error"))
	t.Log(err.(*Errno).PrintStack())
}
