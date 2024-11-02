// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "sfw/biz/handler/api/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_user := _v1.Group("/user", _userMw()...)
				_user.GET("/follower_count", append(_followercountmethodMw(), user.FollowerCountMethod)...)
				_user.GET("/following_count", append(_followingcountmethodMw(), user.FollowingCountMethod)...)
				_user.GET("/info", append(_infomethodMw(), user.InfoMethod)...)
				_user.GET("/like_count", append(_likecountmethodMw(), user.LikeCountMethod)...)
				_user.POST("/login", append(_loginmethodMw(), user.LoginMethod)...)
				_user.POST("/register", append(_registermethodMw(), user.RegisterMethod)...)
				_user.GET("/search", append(_searchmethodMw(), user.SearchMethod)...)
				{
					_avatar := _user.Group("/avatar", _avatarMw()...)
					_avatar.GET("/upload", append(_avataruploadmethodMw(), user.AvatarUploadMethod)...)
				}
				{
					_mfa := _user.Group("/mfa", _mfaMw()...)
					_mfa.POST("/bind", append(_mfabindmethodMw(), user.MfaBindMethod)...)
					_mfa.GET("/qrcode", append(_mfaqrcodemethodMw(), user.MfaQrcodeMethod)...)
				}
				{
					_security := _user.Group("/security", _securityMw()...)
					{
						_email := _security.Group("/email", _emailMw()...)
						_email.POST("/code", append(_securityemailcodemethodMw(), user.SecurityEmailCodeMethod)...)
					}
					{
						_password := _security.Group("/password", _passwordMw()...)
						{
							_reset := _password.Group("/reset", _resetMw()...)
							_reset.POST("/email", append(_passwordresetmethodMw(), user.PasswordResetMethod)...)
						}
						{
							_retrieve := _password.Group("/retrieve", _retrieveMw()...)
							_retrieve.POST("/email", append(_passwordretrievemethodMw(), user.PasswordRetrieveMethod)...)
						}
					}
				}
			}
		}
	}
}
