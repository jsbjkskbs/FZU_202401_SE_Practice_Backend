package generator

import (
	"math/rand"
	"time"
)

type AlnumGeneratorOption struct {
	// Length 长度
	Length int
	// UseLowerAlpha 使用小写字母
	UseLowerAlpha bool
	// UseUpperAlpha 使用大写字母
	UseUpperAlpha bool
	// UseNumber 使用数字
	UseNumber bool
	// UseSpecialChar 使用特殊字符
	UseSpecialChar bool
	// SpecialChar 特殊字符
	SpecialChar string

	// UseCustomChar 使用自定义字符
	UseCustomChar bool
	// CustomChar 自定义字符
	CustomChar string
}

const (
	// LowerAlpha 小写字母
	lowerAlpha = "abcdefghijklmnopqrstuvwxyz"

	// UpperAlpha 大写字母
	upperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Number 数字
	number = "0123456789"

	// SpecialChar 特殊字符
	specialChar = "!@#$%^&*()_+{}|:<>?~"
)

// GenerateAlnumString generates a random string.
// option is the option of the generator.
// GenerateAlnumString 生成一个随机字符串。
// option 是生成器的选项。
func GenerateAlnumString(option AlnumGeneratorOption) string {
	charSet := ""
	if option.UseCustomChar {
		charSet = option.CustomChar
	} else {
		if option.UseLowerAlpha {
			charSet += lowerAlpha
		}
		if option.UseUpperAlpha {
			charSet += upperAlpha
		}
		if option.UseNumber {
			charSet += number
		}
		if option.UseSpecialChar {
			charSet += specialChar
		}
	}
	return generate(charSet, option.Length)
}

// generate generates a random string.
// charSet is the character set.
// length is the length of the string.
// generate 生成一个随机字符串。
// charSet 是字符集。
// length 是字符串的长度。
func generate(charSet string, length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
