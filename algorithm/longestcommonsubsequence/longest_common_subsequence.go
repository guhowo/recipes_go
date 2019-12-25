package longestcommonsubsequence

import "math"

//最长公共子序列，可以不连续
func longestCommonSubsequence(a, b string) int {
	lenA, lenB := len(a), len(b)
	dp := make([][]int, lenA)
	for i, _ := range dp {
		dp[i] = make([]int, lenB)
	}

	//初始化dp第0列和第0行
	for j := 0; j < lenB; j++ {
		if a[0] == b[j] || (j >= 1 && dp[0][j-1] == 1) {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}
	for i := 0; i < lenA; i++ {
		if a[i] == b[0] || (i >= 1 && dp[i-1][0] == 1) {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}

	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[lenA-1][lenB-1]
}
