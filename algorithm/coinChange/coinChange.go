package coinChange

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount <= 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 || dp[i-coin] == -1 {
				continue
			}
			if dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
		if dp[i] == math.MaxInt32 {
			dp[i] = -1
		}
	}
	return dp[amount]
}
