package longestcommonsubsequence

import "testing"

func Test_LongestCommonSubsequence(t *testing.T) {
	for _, unit := range []struct {
		Str1   string
		Str2   string
		Expect int
	}{
		{
			"abcdefg",
			"bchef",
			4,
		},
	} {
		if actual := longestCommonSubsequence(unit.Str1, unit.Str2); actual != unit.Expect {
			t.Errorf("str1 = %s, str2 = %s, expect = %d, actual = %d", unit.Str1, unit.Str2, unit.Expect, actual)
		}
	}
}
