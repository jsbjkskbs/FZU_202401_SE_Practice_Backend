package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type RsaService struct {
	ServerRsaKey *rsa.PrivateKey
	ClientRsaKey *rsa.PublicKey
	Label        []byte
}

func NewRsaService() *RsaService {
	return &RsaService{}
}

func (r *RsaService) Build(clientKey []byte) (err error) {
	r.ServerRsaKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	d := make([]byte, base64.StdEncoding.DecodedLen(len(clientKey)))
	n, err := base64.StdEncoding.Decode(d, clientKey)
	if err != nil {
		return errors.New(`client key base64 decode error`)
	}
	d = d[:n]
	key, err := x509.ParsePKIXPublicKey(d)
	if err != nil {
		return err
	}
	var ok bool
	if r.ClientRsaKey, ok = key.(*rsa.PublicKey); !ok {
		return errors.New(`client key is not rsa public key`)
	}
	return nil
}

func (r *RsaService) BuildWithoutClientKey() (err error) {
	r.ServerRsaKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	r.ClientRsaKey = nil
	return nil
}

func (r *RsaService) Decode(cryptoMsg []byte) (msg []byte, err error) {
	if r.ServerRsaKey == nil {
		return nil, errors.New(`server rsa key is nil`)
	}
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, r.ServerRsaKey, cryptoMsg)
	return
}

func (r *RsaService) Encode(msg []byte) (cryptoMsg []byte, err error) {
	if r.ClientRsaKey == nil {
		return nil, errors.New(`client rsa key is nil`)
	}
	cryptoMsg, err = rsa.EncryptPKCS1v15(rand.Reader, r.ClientRsaKey, msg) // 不用OAEP方便测试
	return
}

func (r *RsaService) Signature(msg []byte) (signature []byte, err error) {
	if r.ServerRsaKey == nil {
		return nil, errors.New(`server rsa key is nil`)
	}
	hashed := sha256.Sum256(msg)
	signature, err = rsa.SignPKCS1v15(rand.Reader, r.ServerRsaKey, crypto.SHA256, hashed[:])
	return
}

func (r *RsaService) Verify(signature []byte, hashed [32]byte) (err error) {
	err = rsa.VerifyPKCS1v15(r.ClientRsaKey, crypto.SHA256, hashed[:], signature)
	return
}

func (r *RsaService) GetPublicKeyPemFormat() (string, error) {
	b, err := x509.MarshalPKIXPublicKey(&r.ServerRsaKey.PublicKey)
	if err != nil {
		return ``, err
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	})
	return string(keyPEM), nil
}

func (r *RsaService) GetPrivateKeyPemFormat() (string, error) {
	b, err := x509.MarshalPKCS8PrivateKey(r.ServerRsaKey)
	if err != nil {
		return ``, err
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	})
	return string(keyPEM), nil
}
