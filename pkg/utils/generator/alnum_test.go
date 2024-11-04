package generator_test

import (
	"testing"

	"sfw/pkg/utils/generator"
)

func TestAlnumAll(t *testing.T) {
	t.Log("TestAlnumAll")
	for i := 0; i < 10; i++ {
		t.Log(generator.GenerateAlnumString(generator.AlnumGeneratorOption{
			Length:         10,
			UseLowerAlpha:  true,
			UseUpperAlpha:  true,
			UseNumber:      true,
			UseSpecialChar: true,
		}))
	}
}

func TestAlnumLower(t *testing.T) {
	t.Log("TestAlnumLower")
	for i := 0; i < 10; i++ {
		t.Log(generator.GenerateAlnumString(generator.AlnumGeneratorOption{
			Length:        10,
			UseLowerAlpha: true,
		}))
	}
}

func TestAlnumUpper(t *testing.T) {
	t.Log("TestAlnumUpper")
	for i := 0; i < 10; i++ {
		t.Log(generator.GenerateAlnumString(generator.AlnumGeneratorOption{
			Length:        10,
			UseUpperAlpha: true,
		}))
	}
}

func TestAlnumNumber(t *testing.T) {
	t.Log("TestAlnumNumber")
	for i := 0; i < 10; i++ {
		t.Log(generator.GenerateAlnumString(generator.AlnumGeneratorOption{
			Length:    10,
			UseNumber: true,
		}))
	}
}

func TestAlnumSpecial(t *testing.T) {
	t.Log("TestAlnumSpecial")
	for i := 0; i < 10; i++ {
		t.Log(generator.GenerateAlnumString(generator.AlnumGeneratorOption{
			Length:         10,
			UseSpecialChar: true,
		}))
	}
}

func TestAlnumCustom(t *testing.T) {
	t.Log("TestAlnumCustom")
	for i := 0; i < 10; i++ {
		t.Log(generator.GenerateAlnumString(generator.AlnumGeneratorOption{
			Length:        10,
			UseCustomChar: true,
			CustomChar:    "abc",
		}))
	}
}
