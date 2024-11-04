package zincsearch

type Logger struct {
	Info    func(v ...any)
	Infof   func(format string, v ...any)
	Error   func(v ...any)
	Errorf  func(format string, v ...any)
	Warn    func(v ...any)
	Warnf   func(format string, v ...any)
	Debug   func(v ...any)
	Debugf  func(format string, v ...any)
	Fatal   func(v ...any)
	Fatalf  func(format string, v ...any)
	Panic   func(v ...any)
	Panicf  func(format string, v ...any)
	Trace   func(v ...any)
	Tracef  func(format string, v ...any)
	Notice  func(v ...any)
	Noticef func(format string, v ...any)
}

func (l *ZincClient) NewLogger(title string) *Logger {
	return &Logger{
		Info:    l.GetDecoInfoFunc(title),
		Infof:   l.GetDecoInfofFunc(title),
		Error:   l.GetDecoErrorFunc(title),
		Errorf:  l.GetDecoErrorfFunc(title),
		Warn:    l.GetDecoWarnFunc(title),
		Warnf:   l.GetDecoWarnfFunc(title),
		Debug:   l.GetDecoDebugFunc(title),
		Debugf:  l.GetDecoDebugfFunc(title),
		Fatal:   l.GetDecoFatalFunc(title),
		Fatalf:  l.GetDecoFatalfFunc(title),
		Panic:   l.GetDecoPanicFunc(title),
		Panicf:  l.GetDecoPanicfFunc(title),
		Trace:   l.GetDecoTraceFunc(title),
		Tracef:  l.GetDecoTracefFunc(title),
		Notice:  l.GetDecoNoticeFunc(title),
		Noticef: l.GetDecoNoticefFunc(title),
	}
}

func (l *ZincClient) NewLoggerWithOtherOutput(
	title string,
	output func(v ...any),
	outputf func(format string, v ...any)) *Logger {
	return &Logger{
		Info:    l.GetDecoInfoFunc(title, output),
		Infof:   l.GetDecoInfofFunc(title, outputf),
		Error:   l.GetDecoErrorFunc(title, output),
		Errorf:  l.GetDecoErrorfFunc(title, outputf),
		Warn:    l.GetDecoWarnFunc(title, output),
		Warnf:   l.GetDecoWarnfFunc(title, outputf),
		Debug:   l.GetDecoDebugFunc(title, output),
		Debugf:  l.GetDecoDebugfFunc(title, outputf),
		Fatal:   l.GetDecoFatalFunc(title, output),
		Fatalf:  l.GetDecoFatalfFunc(title, outputf),
		Panic:   l.GetDecoPanicFunc(title, output),
		Panicf:  l.GetDecoPanicfFunc(title, outputf),
		Trace:   l.GetDecoTraceFunc(title, output),
		Tracef:  l.GetDecoTracefFunc(title, outputf),
		Notice:  l.GetDecoNoticeFunc(title, output),
		Noticef: l.GetDecoNoticefFunc(title, outputf),
	}
}
