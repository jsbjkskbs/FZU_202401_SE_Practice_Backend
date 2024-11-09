package logger

import (
	"sfw/biz/mw/zincsearch"
	"sfw/pkg/errno"
)

var (
	RuntimeLogger     *zincsearch.Logger
	SynchronizeLogger *zincsearch.Logger
	ShutdownLogger    *zincsearch.Logger
)

func InitLogger() {
	RuntimeLogger = zincsearch.Client.NewLogger("runtime")
	SynchronizeLogger = zincsearch.Client.NewLogger("synchronize")
	ShutdownLogger = zincsearch.Client.NewLogger("shutdown")
}

func LogRuntimeError(e error) {
	go logRuntimeError(e)
}

func logRuntimeError(e error) {
	if e == nil {
		return
	}

	err := errno.ConvertErrno(e)
	if err.RecommendToPrintStack() {
		RuntimeLogger.Error(err.PrintStack())
	}
}
