package errno_test

import (
	"errors"
	"testing"

	"sfw/pkg/errno"
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

func TestErrno7(t *testing.T) {
	errno.CustomError.WithMessage("1")
	if errno.CustomError.Message == "1" {
		t.Error("CustomError message should not be '1'")
	}
	errno.CustomError.WithMessage("2")
	if errno.CustomError.Message == "2" {
		t.Error("CustomError message should not be '2'")
	}
}

func TestErrno8(t *testing.T) {
	inner1 := errno.CustomError.InnerErrno
	inner2 := errno.CustomError.WithInnerError(errors.New("123")).InnerErrno
	inner3 := errno.CustomError.InnerErrno
	if !(inner1 == nil && inner1 == inner3) {
		t.Errorf("errno.CustomError has modified, %#v\n", errno.CustomError)
	}
	if inner2 == inner1 || inner2 == inner3 {
		t.Errorf("unexpected new Errno, %#v\n", inner2)
	}
	t.Logf("passed, %#v, %#v, %#v\n", inner1, inner2, inner3)
}
