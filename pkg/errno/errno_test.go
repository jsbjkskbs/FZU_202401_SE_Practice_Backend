package errno_test

import (
	"errors"
	"sfw/pkg/errno"
	"testing"
)

// TestErrno1 test errno
// unsafe operation
func TestErrno1(t *testing.T) {
	var err error
	err = errno.NewErrnoWithInnerErrno(1, "test", errno.InternalServerError)
	t.Log(err.(*errno.Errno).PrintStack())
}

// TestErrno2 test errno
// safe operation
func TestErrno2(t *testing.T) {
	var err error
	err = errors.New("test")
	t.Log(errno.ConvertErrno(err).PrintStack())
}

// TestErrno3 test errno
// safe operation
func TestErrno3(t *testing.T) {
	var err error
	err = errno.NewErrnoWithInnerErrno(1, "test", errno.InternalServerError)
	t.Log(errno.ConvertErrno(err).PrintStack())
}

func TestErrno4(t *testing.T) {
	var err error
	err = errno.NewErrno(1, "test")
	err = err.(*errno.Errno).WithInnerError(errors.New("test inner error"))
	t.Log(err.(*errno.Errno).PrintStack())
}

func TestErrno5(t *testing.T) {
	var err error
	err = errors.New("test")
	err = errno.ConvertErrno(err).WithInnerError(errors.New("test inner error"))
	t.Log(err.(*errno.Errno).PrintStack())
}

func TestErrno6(t *testing.T) {
	var err error
	err = errno.NewErrno(1, "test")
	err = err.(*errno.Errno).WithInnerError(
		errno.NewErrno(2, "test inner error"),
	)
	t.Log(err.(*errno.Errno).PrintStack())
}
