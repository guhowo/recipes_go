package heap

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{-1, 3, 5, 3, 9, 8, 1}
	Sort(a, 1, len(a)-1)
	fmt.Println(a)
}
