package largestsum

// 求一个数组的元素和的最大值，要求所选数字不能相邻

import "math"

// 递归版本O(2^n)
func recOpt(a []float64, i int) float64 {
	if i == 0 {
		return a[0]
	} else if i == 1 {
		return math.Max(a[0], a[1])
	} else {
		A := recOpt(a[:i-1], i-2) + a[i]
		B := recOpt(a[:i], i-1)
		return math.Max(A, B)
	}
}

// 非递归版本
func DpOpt(a []float64, i int) float64 {
	dp := make([]float64, i+1)

	dp[0] = a[0]
	dp[1] = math.Max(a[0], a[1])
	for i := 2; i < len(a); i++ {
		A := dp[i-2] + a[i]
		B := dp[i-1]
		dp[i] = math.Max(A, B)
	}
	return dp[i]
}
