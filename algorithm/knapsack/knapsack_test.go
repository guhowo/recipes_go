package knapsack

import (
	"fmt"
	"testing"
)

func TestBag_Solve(t *testing.T) {
	items := []int{4, 7, 8, 16}
	n := len(items)
	bag := IBag(&Bag{})
	result := bag.Solve(items, n, 30)
	fmt.Println(result)

	bag = IBag(&BagV1{})
	result = bag.Solve(items, n, 30)
	fmt.Println(result)
}
