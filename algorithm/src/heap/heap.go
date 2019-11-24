package heap

func Sort(a []int, start, end int) {
	// 建堆：从下自上
	for parent := end / 2; parent >= start; parent-- {
		MaxHeapify(a, parent, end)
	}

	// 首尾交换
	// 从顶向下调整，保持堆的特性
	for i := 0; i < end-start; i++ {
		a[start], a[end-i] = a[end-i], a[start]
		MaxHeapify(a, start, end-i-1)
	}
}

func MaxHeapify(a []int, start, end int) {
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
