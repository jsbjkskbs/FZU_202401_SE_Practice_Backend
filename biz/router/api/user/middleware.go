// Code generated by hertz generator.

package user

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

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followercountmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followingcountmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _infomethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _likecountmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registermethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _searchmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _avatarMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _avataruploadmethodMw() []app.HandlerFunc {
	// your code...
	return auth.Auth()
}

func _mfaMw() []app.HandlerFunc {
	// your code...
	return auth.Auth()
}

func _mfabindmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _mfaqrcodemethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordresetmethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordretrivemethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _securityMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _emailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _securityemailcodemethodMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordretrievemethodMw() []app.HandlerFunc {
	// your code...
	return nil
}
