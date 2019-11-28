package heap

import "fmt"

func Sort(a []float64, start, end int) {
	// 建堆：从下自上
	BuildHeap(a, start, end)
	// 首尾交换
	// 从顶向下调整，保持堆的特性
	for i := 0; i < end-start; i++ {
		a[start], a[end-i] = a[end-i], a[start]
		MaxHeapify(a, start, end-i-1)
	}
}

// 建堆：从下自上
func BuildHeap(a []float64, start, end int) {
	for parent := end / 2; parent >= start; parent-- {
		MaxHeapify(a, parent, end)
	}
}

// 堆化过程：一组三角数的比较和交换。
// 如果不换，退出；如果交换，则儿子节点也要做响应的堆化调整。
func MaxHeapify(a []float64, start, end int) {
	parent := start
	son := 2 * parent

	for son <= end {
		// 先取两个子节点中较大的为SON
		if son+1 <= end && a[son] < a[son+1] {
			son++
		}
		// 若有交换，向下接着调整，否则parent以下已经全部是堆了
		if a[son] > a[parent] {
			a[son], a[parent] = a[parent], a[son]
			parent = son
			son = 2 * parent
		} else {
			return
		}
	}
}

// 大顶堆的插入
func MaxHeapInsert(a []float64, start, end int, value float64) []float64 {
	a = append(a, value)
	end++
	parent := end / 2
	son := end
	for parent >= start {
		if a[son] > a[parent] {
			a[son], a[parent] = a[parent], a[son]
			son = parent
			parent = parent / 2
		} else {
			return a
		}
	}
	return a
}

// 小顶堆的插入
func MinHeapInsert(a []float64, start, end int, value float64) []float64 {
	a = append(a, value)
	end++
	parent := end / 2
	son := end
	for parent >= start {
		if a[son] < a[parent] {
			son = parent
			parent = parent / 2
		} else {
			return a
		}
	}
	return a
}

func MaxHeapDelete(a []float64, start, end int) []float64 {
	a[start], a[end] = a[end], a[start]
	end--
	parent := start
	son := 2 * parent
	for parent <= end/2 {
		if son+1 <= end && a[son] < a[son+1] {
			son++
		}
		if a[son] > a[parent] {
			a[son], a[parent] = a[parent], a[son]
			parent = son
			son = 2 * parent
		} else {
			return a
		}
	}
	return a[:end+1]
}

func MinHeapDelete(a []float64, start, end int) []float64 {
	a[start], a[end] = a[end], a[start]
	end--
	parent := start
	son := 2 * parent
	for parent <= end/2 {
		if son+1 <= end && a[son] > a[son+1] {
			son++
		}
		if a[son] < a[parent] {
			a[son], a[parent] = a[parent], a[son]
			parent = son
			son = 2 * parent
		} else {
			return a
		}
	}
	return a[:end+1]
}

/*
 * 求数组的中位数
 * 左边是大堆，右边是小堆。
 * 如果插入的数比当前的中位数大，那么插入右边（小堆），反之插入左边（大堆）
 * 每次插入后，两个堆大小不能超过1，否则把多的那个堆顶放到另一个堆
 */
//func Median(a []float64, start, end int) (m float64) {
//	if start == end {
//		return a[start]
//	}
//
//	m = float64(0)
//	// initialize max heap and min heap
//	maxHeap := make([]float64, 0)
//	minHeap := make([]float64, 0)
//	maxHeapLen:=0
//	minHeapLen:=0
//	// top of heap to move from bigger heap to the smaller one
//	var move float64
//	for i, value := range a[start:end] {
//		fmt.Println("i = ", i, ", value = ", value)
//		// flag of even or odd
//		// put into min heap when value is bigger then the current median
//		if value > m {
//			minHeap = MinHeapInsert(minHeap, 0, len(minHeap)-1, value)
//		} else if value < m {
//			// put into max heap when value is less then the current median
//			maxHeap = MaxHeapInsert(maxHeap, 0, len(maxHeap)-1, value)
//		} else {
//			// put into the smaller heap if equal
//			if len(maxHeap) > len(minHeap) {
//				minHeap = MinHeapInsert(minHeap, 0, len(minHeap)-1, value)
//			} else {
//				maxHeap = MaxHeapInsert(maxHeap, 0, len(maxHeap)-1, value)
//			}
//		}
//		fmt.Println("maxHeap = ", maxHeap)
//		fmt.Println("minHeap = ", minHeap)
//		// adjust two heaps
//		if len(maxHeap)-len(minHeap) >= 2 {
//			move = maxHeap[0]
//			maxHeap = MaxHeapDelete(maxHeap, 0, len(maxHeap)-1)
//			maxHeap = maxHeap[]
//			minHeap = MinHeapInsert(minHeap, 0, len(minHeap)-1, move)
//			m = (minHeap[0] + maxHeap[0]) / 2
//		} else if len(minHeap)-len(maxHeap) >= 2 {
//			move = minHeap[0]
//			minHeap = MinHeapDelete(minHeap, 0, len(minHeap)-1)
//			maxHeap = MaxHeapInsert(maxHeap, 0, len(maxHeap)-1, move)
//			m = (minHeap[0] + maxHeap[0]) / 2
//		} else if len(maxHeap)-len(minHeap) == 1 {
//			m = maxHeap[0]
//		} else if len(minHeap)-len(maxHeap) == 1 {
//			m = minHeap[0]
//		} else {
//			m = (minHeap[0] + maxHeap[0]) / 2
//		}
//		fmt.Println("maxHeap = ", maxHeap)
//		fmt.Println("minHeap = ", minHeap)
//		//fmt.Println(i+1, " st = ", m)
//	}
//	return m
//}
func Median(a []float64, start, end int) (m float64) {
	if start == end {
		return a[start]
	}

	m = float64(0)
	// initialize max heap and min heap
	maxHeap := make([]float64, 0)
	minHeap := make([]float64, 0)
	maxHeapLen := 0
	minHeapLen := 0
	// top of heap to move from bigger heap to the smaller one
	var move float64
	for i, value := range a[start:end] {
		fmt.Println("i = ", i, ", value = ", value)
		// flag of even or odd
		// put into min heap when value is bigger then the current median
		if value > m {
			minHeap = MinHeapInsert(minHeap, 1, minHeapLen, value)
			minHeapLen++
		} else if value < m {
			// put into max heap when value is less then the current median
			maxHeap = MaxHeapInsert(maxHeap, 1, maxHeapLen, value)
			maxHeapLen++
		} else {
			// put into the smaller heap if equal
			if maxHeapLen > minHeapLen {
				minHeap = MinHeapInsert(minHeap, 1, minHeapLen, value)
				minHeapLen++
			} else {
				maxHeap = MaxHeapInsert(maxHeap, 1, maxHeapLen, value)
				maxHeapLen++
			}
		}
		fmt.Println("maxHeap = ", maxHeap)
		fmt.Println("minHeap = ", minHeap)
		// adjust two heaps
		if maxHeapLen-minHeapLen >= 2 {
			move = maxHeap[0]
			maxHeap = MaxHeapDelete(maxHeap, 1, maxHeapLen)
			maxHeapLen--
			minHeap = MinHeapInsert(minHeap, 1, minHeapLen, move)
			minHeapLen++
			m = (minHeap[0] + maxHeap[0]) / 2
		} else if minHeapLen-maxHeapLen >= 2 {
			move = minHeap[0]
			minHeap = MinHeapDelete(minHeap, 0, minHeapLen-1)
			maxHeap = MaxHeapInsert(maxHeap, 0, maxHeapLen-1, move)
			m = (minHeap[0] + maxHeap[0]) / 2
		} else if maxHeapLen-minHeapLen == 1 {
			m = maxHeap[0]
		} else if minHeapLen-maxHeapLen == 1 {
			m = minHeap[0]
		} else {
			m = (minHeap[0] + maxHeap[0]) / 2
		}
		fmt.Println("maxHeap = ", maxHeap)
		fmt.Println("minHeap = ", minHeap)
		//fmt.Println(i+1, " st = ", m)
	}
	return m
}
