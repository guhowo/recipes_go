package bag

import (
	"fmt"
	"testing"
)

func TestBag(t *testing.T) {
	items := []int{4, 7, 8, 16}
	n = len(items)
	result = make([]int, n)
	getMaxW(0, 0, 30, items)
	fmt.Println("maxW = ", maxW, "\nresult = ", result)
}
