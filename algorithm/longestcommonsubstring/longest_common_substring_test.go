package longestcommonsubstring

import "testing"

func Test_LongestCommonSubstring(t *testing.T) {
	for _, unit := range []struct {
		Str1   string
		Str2   string
		Expect int
	}{
		{
			"abcdefg",
			"acdhf",
			2,
		}, {
			"123456789",
			"8345697",
			4,
		},
	} {
		if actual := longestCommonSubstring(unit.Str1, unit.Str2); actual != unit.Expect {
			t.Errorf("str1 = %v, str2 = %v, expect = %d, actual = %d", unit.Str1, unit.Str2, unit.Expect, actual)
		}
	}
}
