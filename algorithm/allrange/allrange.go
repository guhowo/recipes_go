package allrange

var result = make([][]int, 0)

// cur表示现在的排列，k表示要交换的位置，m表示长度
func allRange(cur []int, k int, m int) {
	if k == m {
		tmp := make([]int, m)
		copy(tmp, cur)
		result = append(result, tmp)
		return

	}
	/*
		1、把第K个和k+i个一一交换
		2、对交换后的k+1做全排列
		3、交换回来，反复1-3步
	*/
	used := make(map[int]bool)
	for i := k; i < m; i++ {
		if !used[cur[i]] {
			cur[k], cur[i] = cur[i], cur[k]
			allRange(cur, k+1, m)
			cur[k], cur[i] = cur[i], cur[k]
			used[cur[i]] = true
		}

	}

}
