package rmatch

import (
	"testing"
)

func Test_Rmatch(t *testing.T) {
	for _, unit := range []struct {
		Text    string
		Pattern string
		Expect  bool
	}{
		{
			"abcbbaed",
			"*?c*aed",
			true,
		},
		{
			"abcbbaed",
			"?c*aed",
			false,
		},
		{"abcbbaed",
			"*?c*aed",
			true,
		}, {
			"mississippi",
			"mis*is*p*.",
			true,
		}, {
			"aab",
			"c*a*b",
			false,
		},
	} {
		matched = false
		rmatch(0, 0, unit.Text, unit.Pattern)
		if matched != unit.Expect {
			t.Errorf("Expected: [%v], actually: [%v]", unit, matched)
		}
	}
}
