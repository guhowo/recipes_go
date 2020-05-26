package sort

type QuickSort struct {
}

func (s *QuickSort) Do(items []int, n int) {
	qsort(items, 0, n-1)
}

func qsort(items []int, start, end int) {
	if start >= end {
		return
	}
	pos := partition(items, start, end)
	qsort(items, start, pos-1)
	qsort(items, pos+1, end)
}

func partition(items []int, start, end int) (pos int) {
	pivot := items[end]
	//i前面的元素都小于pivot
	i := start
	for j := i; j < end; j++ {
		if items[j] < pivot {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[end] = items[end], items[i]
	return i
}
