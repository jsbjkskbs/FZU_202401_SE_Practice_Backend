package oss

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
)

const (
	avatarUploadCallbackBody = `{
		"key": "$(key)",
		"hash": "$(etag)",
		"fsize": $(fsize),
		"bucket": "$(bucket)",
		"name": "$(x:name)",
		"otype": "avatar",
		"oid": "%v"
	}`
	avatarFopFormat = "imageMogr2/thumbnail/256x256/format/webp/blur/1x0/quality/75|saveas/"
)

func UploadAvatar(filename string, oid int64) (string, error) {
	putPolicy, err := uptoken.NewPutPolicyWithKey(Bucket, "avatar/"+filename, time.Now().Add(1*time.Hour))
	if err != nil {
		return "", err
	}
	putPolicy.SetCallbackUrl(CallbackUrl + "/avatar")
	putPolicy.SetCallbackBody(fmt.Sprintf(avatarUploadCallbackBody, oid))

	saveAvatarEntry := base64.URLEncoding.EncodeToString([]byte(Bucket + ":avatar/" + filename))
	avatarFop := avatarFopFormat + saveAvatarEntry

	persistentOps := strings.Join([]string{avatarFop}, ";")
	pipeline := "avatar_pipeline"
	persistentType := int64(0)
	putPolicy.SetPersistentOps(persistentOps).SetPersistentNotifyUrl(CallbackUrl + "/fop").SetPersistentPipeline(pipeline).SetPersistentType(persistentType)
	upToken, err := uptoken.NewSigner(putPolicy, Mac).GetUpToken(context.Background())
	if err != nil {
		return "", err
	}
	return upToken, nil
}
