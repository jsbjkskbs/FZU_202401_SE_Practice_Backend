package oss

import (
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	Mac              *auth.Credentials
	Cfg              *storage.Config
	OperationManager *storage.OperationManager
)

func Load() {
	Mac = auth.New(AccessKey, SecretKey)
	Cfg = &storage.Config{
		Zone:     &storage.ZoneHuanan,
		UseHTTPS: true,
	}
	OperationManager = storage.NewOperationManager(Mac, Cfg)
}
