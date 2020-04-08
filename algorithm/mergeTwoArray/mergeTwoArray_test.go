package mergeTwoArray

import (
	"fmt"
	"testing"
)

func TestMergeTwoArrays(t *testing.T) {
	a := []int{1, 3, 4, 5}
	b := []int{2, 3, 6, 7, 8, 10}
	res := merge(a, b)
	fmt.Println(res)
}
