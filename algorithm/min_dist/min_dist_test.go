package min_dist

import (
	"testing"
)

func TestTrackBack_MinDist(t *testing.T) {
	m := IMinDist(&DP{})

	for _, unit := range []struct {
		Matrix [][]int
		N      int
		Expect int
	}{
		{
			[][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}},
			4,
			19,
		},
	} {
		m.MinDistDP(unit.Matrix, unit.N)
		if actually := m.MinDistDP(unit.Matrix, unit.N); actually != unit.Expect {
			t.Errorf("unit = %v, actually = %v", unit, actually)
		}
	}
}

func TestDP_MinDistDP(t *testing.T) {
	for _, unit := range []struct {
		Matrix [][]int
		N      int
		Expect int
	}{
		{
			[][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}},
			4,
			19,
		},
	} {
		MinDistDP(unit.N-1, unit.N-1, unit.N, unit.Matrix)
		if actually := MinDistDP(unit.N-1, unit.N-1, unit.N, unit.Matrix); actually != unit.Expect {
			t.Errorf("unit = %v, actually = %v", unit, actually)
		}
	}
}
