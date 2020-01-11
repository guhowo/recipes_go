package expression

import "testing"

func TestCalculate(t *testing.T) {
	for _, unit := range []struct {
		expression string
		Expect     int
	}{
		{
			"3+5*8-6",
			37,
		},
		{
			"1+2+3+4",
			10,
		},
	} {
		if actual := Calculate(unit.expression); actual != unit.Expect {
			t.Errorf("Expression = %s, Expect = %d, actual =%d", unit.expression, unit.Expect, actual)
		}
	}
}
