// Code generated by hertz generator. DO NOT EDIT.

package activity

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	activity "sfw/biz/handler/api/activity"
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
				_activity := _v1.Group("/activity", _activityMw()...)
				_activity.GET("/feed", append(_activityfeedmethodMw(), activity.ActivityFeedMethod)...)
				_activity.GET("/info", append(_activityinfomethodMw(), activity.ActivityInfoMethod)...)
				_activity.GET("/list", append(_activitylistmethodMw(), activity.ActivityListMethod)...)
				_activity.POST("/publish", append(_activitypublishmethodMw(), activity.ActivityPublishMethod)...)
			}
		}
	}
}