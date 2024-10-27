package oss

import "net/http"

func Verify(req *http.Request) (bool, error) {
	return Mac.VerifyCallback(req)
}
