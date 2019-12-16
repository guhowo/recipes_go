package quicksort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	a := []int{10, 5, 14, 1, 2, 12, 14, 15, 3}
	expected := []int{1, 2, 3, 5, 10, 12, 14, 14, 15}

	for _, unit := range []struct {
		input []int
		expected []int
	} {
		{
			a,
			expected,
		},
	} {
		Sort(unit.input, 0, len(unit.input) - 1)
		fmt.Println(unit.input)
		if compare(unit.input, unit.expected) != 0{
			t.Errorf("sort failed")
		}
	}
}


func compare(a []int, b []int) int {
	la := len(a)
	lb := len(b)

	i := 0
	for i < la && i < lb {
		if a[i]>b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
		i++
	}

	if la == lb {
		return 0
	} else if la > lb {
		return 1
	} else {
		return -1
	}
}