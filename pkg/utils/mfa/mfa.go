package mfa

import (
	"github.com/pquerna/otp/totp"
)

// AuthController is a controller for MFA.
// payload is the payload of the MFA.
// code is the code of the MFA.
// secret is the secret of the MFA.
// AuthController 是 MFA 的控制器。
// payload 是 MFA 的载荷。
// code 是 MFA 的代码。
// secret 是 MFA 的密钥。
type AuthController struct {
	payload string
	code    string
	secret  string
}

// MfaAuthInfo represents the information of the MFA.
// Url is the URL of the MFA.
// Secret is the secret of the MFA.
// MfaAuthInfo 表示 MFA 的信息。
// Url 是 MFA 的 URL。
type MfaAuthInfo struct {
	Url    string
	Secret string
}

// NewAuthController creates a new MFA controller.
// payload is the payload of the MFA.
// code is the code of the MFA.
// secret is the secret of the MFA.
// NewAuthController 创建一个新的 MFA 控制器。
// payload 是 MFA 的载荷。
// code 是 MFA 的代码。
// secret 是 MFA 的密钥。
func NewAuthController(payload string, code, secret string) *AuthController {
	return &AuthController{
		payload: payload,
		code:    code,
		secret:  secret,
	}
}

// GenerateTOTP generates a TOTP.
// GenerateTOTP 生成一个 TOTP。
func (ac *AuthController) GenerateTOTP() (*MfaAuthInfo, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "FuliFuli",
		AccountName: ac.payload,
	})

	if err != nil {
		return nil, err
	}

	return &MfaAuthInfo{Url: key.URL(), Secret: key.Secret()}, nil
}

// VerifyTOTP verifies a TOTP.
// VerifyTOTP 验证一个 TOTP。
func (ac *AuthController) VerifyTOTP() bool {
	valid := totp.Validate(ac.code, ac.secret)
	return valid
}
