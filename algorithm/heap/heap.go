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
func MaxHeapInsert(a []float64, start, end int, value float64) (n []float64) {
	if start > end {
		return append(a, value)
	}

	n = append(a, value)
	end++
	parent := end / 2
	son := end
	for parent >= start {
		if n[son] > n[parent] {
			n[son], n[parent] = n[parent], n[son]
			son = parent
			parent = parent / 2
		} else {
			return n
		}
	}
	return n
}

// 小顶堆的插入
func MinHeapInsert(a []float64, start, end int, value float64) (n []float64) {
	if start > end {
		return append(a, value)
	}
	n = append(a, value)
	end++
	parent := end / 2
	son := end
	for parent >= start {
		if n[son] < n[parent] {
			son = parent
			parent = parent / 2
		} else {
			return n
		}
	}
	return n
}

func MaxHeapDelete(a []float64, start, end int) []float64 {
	if start > end {
		return a
	}

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
	if start > end {
		return a
	}

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

func Median(a []float64, start, end int) (m float64) {
	if start == end {
		return a[start]
	}

	m = float64(0)
	// initialize max heap and min heap
	maxHeap := make([]float64, 1)
	minHeap := make([]float64, 1)
	maxHeapLen := 1
	minHeapLen := 1
	// top of heap to move from bigger heap to the smaller one
	var move float64
	for i, value := range a[start:end] {
		fmt.Println("i = ", i, ", value = ", value)
		// flag of even or odd
		// put into min heap when value is bigger then the current median
		if value > m {
			minHeap = MinHeapInsert(minHeap, 1, minHeapLen-1, value)
			minHeapLen++
		} else if value < m {
			// put into max heap when value is less then the current median
			maxHeap = MaxHeapInsert(maxHeap, 1, maxHeapLen-1, value)
			maxHeapLen++
		} else {
			// put into the smaller heap if equal
			if maxHeapLen > minHeapLen {
				minHeap = MinHeapInsert(minHeap, 1, minHeapLen-1, value)
				minHeapLen++
			} else {
				maxHeap = MaxHeapInsert(maxHeap, 1, maxHeapLen-1, value)
				maxHeapLen++
			}
		}

		// adjust two heaps
		if maxHeapLen-minHeapLen >= 2 {
			move = maxHeap[0]
			maxHeap = MaxHeapDelete(maxHeap, 1, maxHeapLen-1)
			maxHeapLen--
			minHeap = MinHeapInsert(minHeap, 1, minHeapLen-1, move)
			minHeapLen++
			m = (minHeap[1] + maxHeap[1]) / 2
		} else if minHeapLen-maxHeapLen >= 2 {
			move = minHeap[1]
			minHeap = MinHeapDelete(minHeap, 1, minHeapLen-1)
			minHeapLen--
			maxHeap = MaxHeapInsert(maxHeap, 1, maxHeapLen-1, move)
			maxHeapLen++
			m = (minHeap[1] + maxHeap[1]) / 2
		} else if maxHeapLen-minHeapLen == 1 {
			m = maxHeap[1]
		} else if minHeapLen-maxHeapLen == 1 {
			m = minHeap[1]
		} else {
			m = (minHeap[1] + maxHeap[1]) / 2
		}
	}
	return m
}
