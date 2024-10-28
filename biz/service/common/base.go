package common

/*

	用于节省代码块，将相同的代码块提取到common.go中

*/

func CorrectPageNumAndPageSize(pageNum, pageSize int64) (int64, int64) {
	return max(pageNum, 0), max(pageSize, 1)
}
