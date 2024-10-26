package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

const (
	// Salt 盐
	salt = "fulifuli"
)

// EncryptBySHA256 SHA256加密
func EncryptBySHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// EncryptBySHA256WithSalt SHA256加密（带盐）
func EncryptBySHA256WithSalt(data, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(data + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetSalt() string {
	return salt
}
