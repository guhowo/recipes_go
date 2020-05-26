package sort

//归并排序
type MergeSort struct {
}

func (m *MergeSort) Do(items []int, n int) {
	mergeSortHelper(items, 0, n-1)
}

func mergeSortHelper(items []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) >> 1
	mergeSortHelper(items, start, mid)
	mergeSortHelper(items, mid+1, end)
	merge(items, start, mid, end)
}

func merge(items []int, start, mid, end int) {
	if start >= end {
		return
	}

	tmp := make([]int, end-start+1)
	k := 0
	i, j := start, mid+1
	for i <= mid && j <= end {
		if items[i] < items[j] {
			tmp[k] = items[i]
			i++
			k++
		} else {
			tmp[k] = items[j]
			j++
			k++
		}
	}

	// 判断哪个子数组中有剩余的数据
	if i <= mid {
		for i <= mid {
			tmp[k] = items[i]
			i++
			k++
		}
	}
	if j <= end {
		for j <= end {
			tmp[k] = items[j]
			j++
			k++
		}
	}

	// 将tmp中的数组拷贝回items[start:]
	copy(items[start:], tmp)
}
