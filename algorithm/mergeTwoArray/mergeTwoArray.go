package mergeTwoArray

import "fmt"

func merge(a, b []int) []int {
	m := len(a)
	n := len(b)
	res := make([]int, m+n)

	i, j := 0, 0
	k := 0
	for i < m && j < n {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
			k++
		} else {
			res[k] = b[j]
			j++
			k++
		}
	}

	fmt.Println(res)

	for i < m {
		res[k] = a[i]
		i++
		k++
	}
	for j < n {
		res[k] = b[j]
		j++
		k++
	}

	return res
}
