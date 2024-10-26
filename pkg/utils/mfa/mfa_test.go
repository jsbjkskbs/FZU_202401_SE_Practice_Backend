package mfa

import (
	"fmt"
	"testing"
)

func TestMfaGenerate(t *testing.T) {
	ac := NewAuthController("test", "test", "test")
	info, err := ac.GenerateTOTP()
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	t.Log("TestMfaGenerate passed, got", fmt.Sprintf("%+v", info))
}
