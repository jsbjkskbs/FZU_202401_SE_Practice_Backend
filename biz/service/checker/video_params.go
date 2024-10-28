package checker

import (
	"sfw/pkg/errno"
	"strings"
)

func CheckVideoPublish(title, desc, category string, labels []string) error {
	if strings.TrimSpace(title) == "" {
		return errno.ParamInvalid
	}
	if strings.TrimSpace(desc) == "" {
		return errno.ParamInvalid
	}
	if strings.TrimSpace(category) == "" {
		return errno.ParamInvalid
	}
	if len(labels) == 0 {
		return errno.ParamInvalid
	}
	if _, ok := CategoryMap[category]; !ok {
		return errno.ParamInvalid
	}
	return nil
}
