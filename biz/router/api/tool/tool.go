// Code generated by hertz generator. DO NOT EDIT.

package tool

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	tool "sfw/biz/handler/api/tool"
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
				_admin := _v1.Group("/admin", _adminMw()...)
				{
					_tool := _admin.Group("/tool", _toolMw()...)
					{
						_delete := _tool.Group("/delete", _deleteMw()...)
						_delete.DELETE("/activity", append(_admintooldeleteactivityMw(), tool.AdminToolDeleteActivity)...)
						_delete.DELETE("/comment", append(_admintooldeletecommentMw(), tool.AdminToolDeleteComment)...)
						_delete.DELETE("/video", append(_admintooldeletevideoMw(), tool.AdminToolDeleteVideo)...)
					}
				}
			}
			{
				_tool0 := _v1.Group("/tool", _tool0Mw()...)
				{
					_delete0 := _tool0.Group("/delete", _delete0Mw()...)
					_delete0.DELETE("/activity", append(_tooldeleteactivityMw(), tool.ToolDeleteActivity)...)
					_delete0.DELETE("/comment", append(_tooldeletecommentMw(), tool.ToolDeleteComment)...)
					_delete0.DELETE("/video", append(_tooldeletevideoMw(), tool.ToolDeleteVideo)...)
				}
				{
					_get := _tool0.Group("/get", _getMw()...)
					_get.GET("/image", append(_toolgetimageMw(), tool.ToolGetImage)...)
				}
				{
					_refresh_token := _tool0.Group("/refresh_token", _refresh_tokenMw()...)
					_refresh_token.GET("/refresh", append(_toolrefreshtokenrefreshMw(), tool.ToolRefreshTokenRefresh)...)
				}
				{
					_token := _tool0.Group("/token", _tokenMw()...)
					_token.GET("/refresh", append(_tooltokenrefreshMw(), tool.ToolTokenRefresh)...)
				}
				{
					_upload := _tool0.Group("/upload", _uploadMw()...)
					_upload.GET("/image", append(_tooluploadimageMw(), tool.ToolUploadImage)...)
				}
			}
		}
	}
}
