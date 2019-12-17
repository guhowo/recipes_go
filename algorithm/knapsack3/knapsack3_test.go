package knapsack3

import (
	"testing"
)

func TestBag_Solve(t *testing.T) {
	bag := IBag(&Bag{})
	for _, unit := range []struct {
		Weight []int
		Values []int
		W      int
		Expect int
	}{
		{
			[]int{2, 2, 4, 6, 3},
			[]int{3, 4, 8, 9, 6},
			9,
			18,
		},
	} {
		if actually := bag.Solve(unit.Weight, unit.Values, len(unit.Weight), unit.W); actually != unit.Expect {
			t.Errorf("case = %v, actually = %v", unit, actually)
		}
	}
}
