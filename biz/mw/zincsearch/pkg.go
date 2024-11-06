package zincsearch

import (
	"fmt"
	"runtime"

	zincsearch "github.com/zinclabs/sdk-go-zincsearch"
)

func (l *ZincClient) Ping() error {
	_, _, err := l.client.Default.Healthz(l.ctx).Execute()
	return err
}

// https://zincsearch-docs.zinc.dev/api/index/exists/
// Check index exists
// Endpoint - HEAD /api/index/:target
// Request
// e.g. HEAD http://localhost:4080/api/index/myindex
// It will response http code 200 or 404.
// 200 means the index exists.
// 404 means the index does not exist.
// if err != nil, means the index does not exist instead of other errors.
func (l *ZincClient) indexExists(index string) bool {
	_, _, err := l.client.Index.Exists(l.ctx, index).Execute()
	if err != nil {
		return false
	}
	return true
}

// TryInitialize initializes the logger.
// It returns a boolean and an error.
// The boolean indicates whether the logger is already initialized.
// The error indicates if the request fails.
// If false, means that it's not initialized yet.
// If true, means that it's already initialized.
// It is used to check if the logger is already initialized.
func (l *ZincClient) TryInitialize() (bool, error) {
	if l.indexExists("log") {
		return true, nil
	}

	data := zincsearch.NewMetaIndexSimple()
	data.SetName("log")
	data.SetMappings(GetDefaultProperties())
	data.SetShardNum(1)
	data.SetStorageType("disk")

	_, _, err := l.client.Index.Create(l.ctx).Data(*data).Execute()
	if err != nil {
		return false, err
	}

	return true, nil
}

// Info logs the data with the status "Info".
// It returns an error if the request fails.
func (l *ZincClient) info(v *Document) error {
	return l.custom(v, "Info")
}

// Error logs the data with the status "Error".
// It returns an error if the request fails.
func (l *ZincClient) error(v *Document) error {
	return l.custom(v, "Error")
}

// Warn logs the data with the status "Warn".
// It returns an error if the request fails.
func (l *ZincClient) warn(v *Document) error {
	return l.custom(v, "Warn")
}

// Debug logs the data with the status "Debug".
// It returns an error if the request fails.
func (l *ZincClient) debug(v *Document) error {
	return l.custom(v, "Debug")
}

// Fatal logs the data with the status "Fatal".
// It returns an error if the request fails.
func (l *ZincClient) fatal(v *Document) error {
	return l.custom(v, "Fatal")
}

// Panic logs the data with the status "Panic".
// It returns an error if the request fails.
func (l *ZincClient) panic(v *Document) error {
	return l.custom(v, "Panic")
}

// Trace logs the data with the status "Trace".
// It returns an error if the request fails.
func (l *ZincClient) trace(v *Document) error {
	return l.custom(v, "Trace")
}

func (l *ZincClient) notice(v *Document) error {
	return l.custom(v, "Notice")
}

// Custom logs the data with the custom status.
// It returns an error if the request fails.
func (l *ZincClient) custom(v *Document, status string) error {
	_, _, err := l.client.Document.
		Index(l.ctx, "log").
		Document(kvDocument(v, l.getCaller(), status)).
		Execute()
	return err
}

// getCaller returns the caller of the function.
// It returns the name and the location of the caller.
// It is used to log the caller of the function.
// Automatically invoked by the logging functions.
func (l *ZincClient) getCaller() _Caller {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return _Caller{}
	}
	return _Caller{
		Location: fmt.Sprintf("%s:%d", file, line),
	}
}
