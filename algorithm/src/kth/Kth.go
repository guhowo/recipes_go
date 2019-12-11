package kth

func getPosition(a []int, left, right int) ([]int, int) {
	if left >= right {
		return a, left
	}

	key := a[left]
	i, j := left, right
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

	return a, i
}

// k从0开始，比如查找第0大的数，此时k==0
func getKth(a []int, k int) int {
	if k >= len(a) || k < 0 {
		return -1
	}

	i, j := 0, len(a)-1
	a, pos := getPosition(a, i, j)
	for pos != k {
		if pos > k {
			a, pos = getPosition(a, i, pos-1)
		} else if pos < k {
			a, pos = getPosition(a, pos+1, j)
		} else {
			return a[k]
		}
	}

	return k
}
