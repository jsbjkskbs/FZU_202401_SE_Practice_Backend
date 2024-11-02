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

	videoUploadCallbackBody = `{
		"key": "$(key)",
		"hash": "$(etag)",
		"fsize": $(fsize),
		"bucket": "$(bucket)",
		"name": "$(x:name)",
		"otype": "video",
		"oid": "%v"
	}`
	// videoFopFormat  = "avthumb/mp4/vcodec/libx264/vb/1.25m/r/24/s/x720|saveas/"

	coverUploadCallbackBody = `{
		"key": "$(key)",
		"hash": "$(etag)",
		"fsize": $(fsize),
		"bucket": "$(bucket)",
		"name": "$(x:name)",
		"otype": "cover",
		"oid": "%v"
	}`
	// coverFopFormat = "imageMogr2/thumbnail/160000@/format/jpg/blur/1x0/quality/75|saveas/"

	AvatarUploadTokenDeadline = 1 * time.Hour
	VideoUploadTokenDeadline  = 1 * time.Hour
	CoverUploadTokenDeadline  = 1 * time.Hour
)

func UploadAvatar(filename string, oid int64) (string, string, error) {
	putPolicy, err := uptoken.NewPutPolicyWithKey(Bucket, "avatar/"+filename, time.Now().Add(AvatarUploadTokenDeadline))
	if err != nil {
		return "", "", err
	}
	putPolicy.SetCallbackUrl(CallbackUrl + "/avatar")
	putPolicy.SetCallbackBody(fmt.Sprintf(avatarUploadCallbackBody, oid))

	saveAvatarEntry := base64.URLEncoding.EncodeToString([]byte(Bucket + ":avatar/" + filename))
	avatarFop := avatarFopFormat + saveAvatarEntry

	persistentOps := strings.Join([]string{avatarFop}, ";")
	persistentType := int64(0)
	putPolicy.SetPersistentOps(persistentOps).SetPersistentNotifyUrl(CallbackUrl + "/fop").SetPersistentType(persistentType)
	upToken, err := uptoken.NewSigner(putPolicy, Mac).GetUpToken(context.Background())
	if err != nil {
		return "", "", err
	}
	return upToken, "avatar/" + filename, nil
}

func UploadVideo(filename string, oid int64) (string, string, error) {
	putPolicy, err := uptoken.NewPutPolicyWithKey(Bucket, "video/"+filename+"/video.mp4", time.Now().Add(VideoUploadTokenDeadline))
	if err != nil {
		return "", "", err
	}
	putPolicy.SetCallbackUrl(CallbackUrl + "/video")
	putPolicy.SetCallbackBody(fmt.Sprintf(videoUploadCallbackBody, oid))

	// saveVideoEntry := base64.URLEncoding.EncodeToString([]byte(Bucket + ":video/" + filename + "/video.mp4"))
	// videoFop := videoFopFormat + saveVideoEntry

	// persistentOps := strings.Join([]string{videoFop}, ";")
	// persistentType := int64(0)
	// putPolicy.SetPersistentOps(persistentOps).SetPersistentNotifyUrl(CallbackUrl + "/fop").SetPersistentType(persistentType)
	upToken, err := uptoken.NewSigner(putPolicy, Mac).GetUpToken(context.Background())
	if err != nil {
		return "", "", err
	}
	return upToken, "video/" + filename + "/video.mp4", nil
}

func UploadVideoCover(filename string, oid int64) (string, string, error) {
	putPolicy, err := uptoken.NewPutPolicyWithKey(Bucket, "video/"+filename+"/cover.jpg", time.Now().Add(CoverUploadTokenDeadline))
	if err != nil {
		return "", "", err
	}
	putPolicy.SetCallbackUrl(CallbackUrl + "/cover")
	putPolicy.SetCallbackBody(fmt.Sprintf(coverUploadCallbackBody, oid))

	// saveVideoEntry := base64.URLEncoding.EncodeToString([]byte(Bucket + ":video/" + filename + "/cover.jpg"))
	// coverFop := coverFopFormat + saveVideoEntry

	// persistentOps := strings.Join([]string{coverFop}, ";")
	// persistentType := int64(0)
	// putPolicy.SetPersistentOps(persistentOps).SetPersistentNotifyUrl(CallbackUrl + "/fop").SetPersistentType(persistentType)
	upToken, err := uptoken.NewSigner(putPolicy, Mac).GetUpToken(context.Background())
	if err != nil {
		return "", "", err
	}
	return upToken, "video/" + filename + "/cover.jpg", nil
}
