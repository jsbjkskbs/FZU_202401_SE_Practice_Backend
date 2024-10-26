package convert

const (
	table = `RNYVWKdGM1xOHovIbnuefPTr5L7X0a9AywqlkCEsgFZSht3B4QU6zic2jpDJ8m`
	xor   = 0x1145141919810
	add   = 0xfc11fc11666
)

// FvEncode convert video id to 13-bit string
// Video ID is an int64 number, convert it to a 13-bit string
// FvEncode 视频ID转换
// 视频ID转换为13位字符串
// 视频ID是一个int64类型的数字，转换为13位字符串
func FvEncode(id int64) string {
	div := id ^ xor + add
	result := ""
	for i := 0; i < 13; i++ {
		index := (div >> (i * 5)) & 0x1F
		result += string(table[index%int64(len(table))])
	}
	return result
}

// FvDecode convert 13-bit string to video id
// FvDecode 13位字符串转换为视频ID
func FvDecode(fv string) int64 {
	div := int64(0)
	for i := 0; i < 13; i++ {
		index := int64(0)
		for j := 0; j < len(table); j++ {
			if table[j] == fv[i] {
				index = int64(j)
				break
			}
		}
		div += index << (i * 5)
	}
	return (div - add) ^ xor
}
