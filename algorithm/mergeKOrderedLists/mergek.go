package mergeKOrderedLists

import (
	"container/heap"
	"fmt"
)

type element struct {
	data     int
	idx      int
	whichArr int
}

type myHeap []element

/* 实现排序 */
func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i] //整个数据交换，包括索引
}

// 最小堆实现
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].data < (*h)[j].data
}

/* 实现往堆中添加元素 */
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(element))
}

/* 实现删除堆中元素 */
func (h *myHeap) Pop() (v interface{}) {
	//old := *h
	//n := len(old)
	//x := old[0]
	//*h = old[1:n]
	//return x
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

// 合并K个有序数组
func MergeKSortedArray(nums [][]int, k int) []int {
	aHeap := &myHeap{}
	// 堆排序处理
	heap.Init(aHeap)
	// 插入nums[x][0]
	for i := 0; i < k; i++ {
		ele := element{data: nums[i][0], idx: 0, whichArr: i}
		heap.Push(aHeap, ele)
		fmt.Println(*aHeap)
	}

	res := []int{}

	for aHeap.Len() > 0 {
		curMaxEle := heap.Pop(aHeap).(element)
		res = append(res, curMaxEle.data)
		if curMaxEle.idx < len(nums[curMaxEle.whichArr])-1 {
			heap.Push(aHeap, element{data: nums[curMaxEle.whichArr][curMaxEle.idx+1], idx: curMaxEle.idx + 1, whichArr: curMaxEle.whichArr})
		}
	}
	return res
}
