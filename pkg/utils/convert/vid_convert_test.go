package convert_test

import (
	"testing"

	"sfw/pkg/utils/convert"
)

func TestVideoIDConvert(t *testing.T) {
	i := int64(234567890000)
	s := convert.FvEncode(i)
	t.Log(s)
	t.Log(convert.FvDecode(s))
	if i != convert.FvDecode(s) {
		t.Error("Expected", i, "got", convert.FvDecode(s))
	} else {
		t.Log("TestVideoIDConvert passed")
	}
}
