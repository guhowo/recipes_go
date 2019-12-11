package merge_sort

type ISort interface {
	Sort(a []int)
}

type Merge struct{}

func (m *Merge) Sort(a []int) {
	m.MergeSort(0, len(a)-1, a)
}

func (m *Merge) MergeSort(start, end int, a []int) {
	if start >= end {
		return
	}

	mid := start + (end-start)>>1
	m.MergeSort(start, mid, a)
	m.MergeSort(mid+1, end, a)
	m.merge(start, end, mid, a)

}

func (m *Merge) merge(low, high, mid int, a []int) {
	temp := make([]int, high-low+1)
	i := low
	j := mid + 1
	k := 0
	for i <= mid && j <= high {
		if a[i] < a[j] {
			temp[k] = a[i]
			k++
			i++
		} else {
			temp[k] = a[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = a[i]
		k++
		i++
	}
	for j <= high {
		temp[k] = a[j]
		k++
		j++
	}

	idx := low
	for i := 0; i <= high-low; i++ {
		a[idx] = temp[i]
		idx++
	}
}
