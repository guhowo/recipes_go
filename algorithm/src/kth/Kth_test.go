package kth

import (
	"fmt"
	"testing"
)

func Test_GetKth(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	k := getKth(a, 0)
	fmt.Println(k)

	k = getKth(a, 1)
	fmt.Println(k)

	k = getKth(a, 2)
	fmt.Println(k)

	k = getKth(a, 3)
	fmt.Println(k)

	k = getKth(a, 4)
	fmt.Println(k)

	k = getKth(a, -1)
	fmt.Println(k)
}
