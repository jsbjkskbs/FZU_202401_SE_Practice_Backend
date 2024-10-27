package oss

import (
	"time"

	"github.com/qiniu/go-sdk/v7/storage"
)

func Key2Url(key string) string {
	deadline := time.Now().Add(1 * time.Hour).Unix()
	url := storage.MakePrivateURLv2(Mac, Domain, key, deadline)
	return url
}
