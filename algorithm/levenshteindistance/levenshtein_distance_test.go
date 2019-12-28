package levenshteindistance

import "testing"

func Test_LevenshteinDistance(t *testing.T) {
	for _, unit := range []struct {
		str1   string
		str2   string
		Expect int
	}{
		{
			"a",
			"ab",
			1,
		},
		{
			"horse",
			"ros",
			3,
		},
		{
			"pneumonoultramicroscopicsilicovolcanoconiosis",
			"ultramicroscopically",
			27,
		},
	} {
		if actual := LevenshteinDistance(unit.str1, unit.str2); actual != unit.Expect {
			t.Errorf("str1 =%s, str2 = %s, expect = %d, actual = %d", unit.str1, unit.str2, unit.Expect, actual)
		}
	}
}
