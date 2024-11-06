package zincsearch

import (
	"context"

	zincsearch "github.com/zinclabs/sdk-go-zincsearch"
)

// ZincClientOption is a struct that contains the host, username, and password.
// It is used to initialize the logger.
// It has the following fields:
// - Host(don't include http:// or https:// and don't forget the port)
// - Username
// - Password
type ZincClientOption struct {
	Host     string
	Username string
	Password string
}

// ZincClient is a struct that contains the context and the client.
// It is used to log the data to the zincsearch.
// It has the following methods:
// - Info
// - Error
// - Warn
// - Debug
// - Fatal
// - Panic
// - Trace
// - Notice
// - Custom
type ZincClient struct {
	ctx    context.Context
	client *zincsearch.APIClient
}

// NewZincClient initializes the logger with the given options.
// It returns the logger.
func NewZincClient(opt *ZincClientOption) *ZincClient {
	ctx := context.WithValue(context.Background(), zincsearch.ContextBasicAuth, zincsearch.BasicAuth{
		UserName: opt.Username,
		Password: opt.Password,
	})
	configuration := zincsearch.NewConfiguration()
	configuration.Host = opt.Host
	client := zincsearch.NewAPIClient(configuration)
	return &ZincClient{
		ctx:    ctx,
		client: client,
	}
}
