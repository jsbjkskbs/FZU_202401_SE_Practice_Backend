// Code generated by hertz generator.

package report

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

func _reportMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminMw() []app.HandlerFunc {
	// your code...
	return auth.AdminAuth()
}

func _report0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminreporthandleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminreportlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _videoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminvideohandleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminvideolistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _reportactivityMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _reportcommentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _reportvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _activityMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminactivityreporthandleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminactivityreportlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _admincommenthandleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _admincommentreportlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _report1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminvideoreporthandleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminvideoreportlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _report2Mw() []app.HandlerFunc {
	// your code...
	return nil
}
