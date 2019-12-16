package largestsum

import (
	"testing"
)

type Input struct {
	arr []float64
	pos int
}

func Test_RecOpt(t *testing.T) {
	for _, unit := range []struct {
		input    Input
		expected float64
	}{
		{
			Input{
				[]float64{1, 2, 4, 1, 7, 8, 3},
				6,
			},
			15,
		},
	} {
		if recOpt(unit.input.arr, unit.input.pos) != unit.expected {
			t.Errorf("test failed")
		}
	}
}

func Test_DpOpt(t *testing.T) {
	for _, unit := range []struct {
		input    Input
		expected float64
	}{
		{
			Input{
				[]float64{1, 2, 4, 1, 7, 8, 3},
				6,
			},
			15,
		},
	} {
		if DpOpt(unit.input.arr, unit.input.pos) != unit.expected {
			t.Errorf("test failed")
		}
	}
}
