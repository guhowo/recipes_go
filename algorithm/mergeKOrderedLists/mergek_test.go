package mergeKOrderedLists

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMergeKSortedArray(t *testing.T) {
	nums := [][]int{}
	nums = append(nums, []int{1, 5, 8})
	nums = append(nums, []int{4, 6, 7})
	nums = append(nums, []int{2, 3, 9})

	fmt.Println(MergeKSortedArray(nums, 3))
}

func TestMyHeap(t *testing.T) {
	//nums := []int{1, 4, 5, 2, 9, 8, 10, 3, 6}
	nums := &myHeap{}
	heap.Init(nums)
	heap.Push(nums, element{data: 6})
	heap.Push(nums, element{data: 1})

	for nums.Len() > 0 {
		v := heap.Pop(nums).(element)
		fmt.Println(v.data)
	}

}
