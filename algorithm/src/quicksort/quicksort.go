package quicksort

func Sort(a []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	key := a[i]

	for i < j {
		for i < j && a[j] >= key {
			j--
		}
		a[i] = a[j]
		for i < j && a[i] <= key {
			i++
		}
		a[j] = a[i]
	}

	a[i] = key
	Sort(a, left, i - 1)
	Sort(a, i + 1, right)
}