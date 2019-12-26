package longestcommonsubstring

import "math"

//最长公共子串，一定要连续
func longestCommonSubstring(a, b string) int {
	dp := make([][]int, len(a))
	for i, _ := range dp {
		dp[i] = make([]int, len(b))
	}

	//初始化第一列
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	//初始化第一行
	for j := 0; j < len(b); j++ {
		if a[0] == b[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	longest := 0
	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				longest = int(math.Max(float64(dp[i][j]), float64(longest)))
			} else {
				dp[i][j] = 0
			}
		}
	}

	return longest
}
