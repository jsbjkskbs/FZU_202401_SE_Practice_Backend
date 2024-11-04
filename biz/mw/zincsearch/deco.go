package zincsearch

import "fmt"

func (l *ZincClient) GetDecoInfoFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Info", outputFunc...)
}

func (l *ZincClient) GetDecoInfofFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Info", outputFunc...)
}

func (l *ZincClient) GetDecoErrorFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Error", outputFunc...)
}

func (l *ZincClient) GetDecoErrorfFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Error", outputFunc...)
}

func (l *ZincClient) GetDecoWarnFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Warn", outputFunc...)
}

func (l *ZincClient) GetDecoWarnfFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Warn", outputFunc...)
}

func (l *ZincClient) GetDecoDebugFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Debug", outputFunc...)
}

func (l *ZincClient) GetDecoDebugfFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Debug", outputFunc...)
}

func (l *ZincClient) GetDecoTraceFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Trace", outputFunc...)
}

func (l *ZincClient) GetDecoTracefFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Trace", outputFunc...)
}

func (l *ZincClient) GetDecoFatalFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Fatal", outputFunc...)
}

func (l *ZincClient) GetDecoFatalfFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Fatal", outputFunc...)
}

func (l *ZincClient) GetDecoPanicFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Panic", outputFunc...)
}

func (l *ZincClient) GetDecoPanicfFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Panic", outputFunc...)
}

func (l *ZincClient) GetDecoNoticeFunc(title string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return l.GetDecoCustomFunc(title, "Notice", outputFunc...)
}

func (l *ZincClient) GetDecoNoticefFunc(title string,
	outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return l.GetDecoCustomfFunc(title, "Notice", outputFunc...)
}

func (l *ZincClient) GetDecoCustomFunc(title string, custom string,
	outputFunc ...func(v ...any)) func(v ...any) {
	return func(v ...any) {
		l.custom(&Document{Title: title, Content: fmt.Sprint(v...)}, custom)
		if len(outputFunc) > 0 {
			for _, f := range outputFunc {
				f(v...)
			}
		}
	}
}

func (l *ZincClient) GetDecoCustomfFunc(title string, custom string, outputFunc ...func(format string, v ...any)) func(format string, v ...any) {
	return func(format string, v ...any) {
		l.custom(&Document{Title: title, Content: fmt.Sprintf(format, v...)}, custom)
		if len(outputFunc) > 0 {
			for _, f := range outputFunc {
				f(format, v...)
			}
		}
	}
}
