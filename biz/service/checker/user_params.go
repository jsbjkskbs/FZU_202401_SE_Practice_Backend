package checker

import (
	"errors"
	"strings"

	"github.com/dlclark/regexp2"
)

func CheckUsername(username string) error {
	if len(username) == 0 {
		return errors.New("用户名不能为空")
	}

	if strings.ContainsAny(username, " \t\n") {
		return errors.New("用户名不能包含空白符")
	}

	return nil
}

var (
	passwordReg = []*regexp2.Regexp{
		regexp2.MustCompile(`(?=.*[0-9])(?=.*[a-zA-Z])(?=.*[^a-zA-Z0-9]).{8,}`, regexp2.None),
		regexp2.MustCompile(`(?=.*[0-9])(?=.*[a-zA-Z]).{8,}`, regexp2.None),
		regexp2.MustCompile(`(?=.*[0-9])(?=.*[^a-zA-Z0-9]).{8,}`, regexp2.None),
		regexp2.MustCompile(`(?=.*[a-zA-Z])(?=.*[^a-zA-Z0-9]).{8,}`, regexp2.None),
	}
)

func CheckPassword(password string) error {
	if len(password) == 0 {
		return errors.New("密码不能为空")
	}

	if strings.ContainsAny(password, " \t\n") {
		return errors.New("密码不能包含空白符")
	}

	ok := false
	for _, reg := range passwordReg {
		if match, _ := reg.MatchString(password); match {
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("密码必须包含字母、数字和特殊字符其中两种, 长度至少8位")
	}

	return nil
}
