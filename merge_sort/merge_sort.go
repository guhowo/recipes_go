package merge_sort

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)>>1
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	mergeTwoSort(a, start, end, mid)
}

func mergeTwoSort(a []int, start, end, mid int) {
	result := make([]int, end-start+1)

	i, j := start, mid+1
	k := 0
	for i <= mid && j <= end {
		if a[i] < a[j] {
			result[k] = a[i]
			i++
			k++
		} else {
			result[k] = a[j]
			j++
			k++
		}
	}

	//if i <= mid {
	//	result = append(result, a[i:mid+1]...)
	//}
	//if j <= end {
	//	result = append(result, a[j:end+1]...)
	//}

	for i <= mid {
		result[k] = a[i]
		k++
		i++
	}
	for j <= end {
		result[k] = a[j]
		k++
		j++
	}

	idx := start
	for k := 0; idx <= end; k++ {
		a[idx] = result[k]
		idx++
	}
}
