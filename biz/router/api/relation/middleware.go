// Code generated by hertz generator.

package relation

import (
	auth "sfw/biz/router/api/api_auth"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followactionmethodMw() []app.HandlerFunc {
	// your code...
	return auth.Auth()
}

func _followlistmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followedMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followedlistmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendlistmethodMw() []app.HandlerFunc {
	// your code...
	return auth.Auth()
}

func _followerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerlistmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}
