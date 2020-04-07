package findMedianSortedArrays

import "testing"

func Test_FindMedianSortedArrays(t *testing.T) {
	for _, unit := range []struct {
		Num1   []int
		Num2   []int
		Expect float64
	}{
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1}, []int{}, 1.0},
		{[]int{1, 3}, []int{2}, 2.0},
	} {
		if actual := findMedianSortedArrays(unit.Num1, unit.Num2); actual != unit.Expect {
			t.Errorf("failed")
		}
	}
}
