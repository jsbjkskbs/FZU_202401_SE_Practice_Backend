package encrypt

import (
	"testing"
)

func TestBase64(t *testing.T) {
	str := "Hello World"
	encoded := EncodeStringToBase64(str)
	decoded, err := DecodeBase64ToString(encoded)
	if err != nil {
		t.Error(err)
	}
	if str != decoded {
		t.Errorf("Expected %s, got %s", str, decoded)
	}
}

func TestQrcode(t *testing.T) {
	url := "https://www.google.com"
	qr := EncodeUrlToQrcodeAsPng(url)
	if qr == "" {
		t.Error("Expected a qrcode, got empty string")
	}
	t.Log("got qr:", qr)
	t.Log("you can verify the qrcode by third party tools")
}
