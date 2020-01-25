package recursive

//数据的全排列
func permute(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)
	allRange(0, n, nums, &result)
	return result
}

//k:起点
//n:长度
//cur:当前排列
func allRange(k, n int, cur []int, result *[][]int) {
	if k == n {
		tmp := make([]int, n)
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for i := k; i < n; i++ {
		cur[i], cur[k] = cur[k], cur[i]
		allRange(k+1, n, cur, result)
		cur[i], cur[k] = cur[k], cur[i]
	}
}
