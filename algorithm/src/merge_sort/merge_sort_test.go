package merge_sort

import (
	"fmt"
	"testing"
)

func TestMerge_Sort(t *testing.T) {
	a := []int{2, 1, 3, 4, 6, 5}
	merge := ISort(&Merge{})
	merge.Sort(a)
	fmt.Println(a)

}
