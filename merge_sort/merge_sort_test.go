package merge_sort

import (
	"fmt"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	a := []int{4, 6, 2, 3, 1, 2, 9}
	mergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}
