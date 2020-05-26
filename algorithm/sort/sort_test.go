package sort

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	var handler ISort

	//归并排序
	handler = &MergeSort{}
	items := []int{11, 8, 3, 9, 7, 1, 2, 5}
	handler.Do(items, len(items))
	fmt.Println("merge sort: ", items)

	//快排
	handler = &QuickSort{}
	items = []int{11, 8, 3, 9, 7, 1, 2, 5}
	handler.Do(items, len(items))
	fmt.Println("quick sort: ", items)
}
