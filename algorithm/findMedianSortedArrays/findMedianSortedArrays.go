package findMedianSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i := 0
	j := 0
	step := (m+n)/2 + 1
	tmp := make([]int, step)
	k := 0
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			tmp[k] = nums2[j]
			j++
			k++
		} else {
			tmp[k] = nums1[i]
			i++
			k++
		}
		if k == step {
			break
		}
	}
	for i < m && k < step {
		tmp[k] = nums1[i]
		k++
		i++
	}
	for j < n && k < step {
		tmp[k] = nums2[j]
		k++
		j++
	}

	if (m+n)%2 == 1 {
		return float64(tmp[len(tmp)-1])
	}

	return float64(tmp[len(tmp)-1]+tmp[len(tmp)-2]) / 2
}
