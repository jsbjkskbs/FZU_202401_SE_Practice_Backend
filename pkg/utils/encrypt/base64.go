package encrypt

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/skip2/go-qrcode"
)

func EncodeUrlToQrcodeAsPng(url string) string {
	code, _ := qrcode.New(url, qrcode.Low)
	img := code.Image(256)
	buf := bytes.NewBuffer(make([]byte, 0))
	png.Encode(buf, img)
	return `data:image/png;base64,` + base64.StdEncoding.EncodeToString(buf.Bytes())
}

func DecodeQrcodeToUrl(qrcode string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(qrcode)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func EncodeStringToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeBase64ToString(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
