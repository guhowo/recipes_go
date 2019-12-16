package fibonacci

import "testing"

func TestSort(t *testing.T) {
	for _, unit := range []struct {
		input    int
		expected int
	}{
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
	} {
		fib := Fib(unit.input)
		if fib != unit.expected {
			t.Error("input = ", unit.input, ", expected = ", unit.expected, ", result = ", fib)
		}
	}
}
