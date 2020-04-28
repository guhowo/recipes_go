package longestvalid

import "fmt"

func longest(s string) int {
	stack := make([]int, 0)

	for idx, c := range s {
		//如果是左括号，压栈
		if c == '(' {
			stack = append(stack, idx)
		}
		//如果是右括号，尝试出栈。如果出栈失败，则入栈
		if c == ')' {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if s[top] == '(' {
					if len(stack) >= 1 {
						stack = stack[:len(stack)-1]
					}
				} else {
					stack = append(stack, idx)
				}
			} else {
				stack = append(stack, idx)
			}
		}
	}
	stack = append(stack, len(s))

	maxLen := 0
	if len(stack) >= 1 {
		maxLen = stack[0]
	}
	for i := 1; i < len(stack); i++ {
		curLen := stack[i] - stack[i-1] - 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	fmt.Println(stack)

	return maxLen
}

func dpLongest(s string) int {
	maxLen := 0
	dp := make([]int, len(s))
	if len(dp) > 0 {
		dp[0] = 0
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		}
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] += dp[i-2]
				}
			}
			if s[i-1] == ')' {
				if dp[i-1] == 0 {
					dp[i] = 0
				} else if i-dp[i-1]-1 >= 0 {
					if dp[i-dp[i-1]-1] == '(' {
						dp[i] = dp[i-1] + 2
						if i-dp[i-1]-2 >= 0 {
							dp[i] += dp[i-dp[i-1]-2]
						}
					} else {
						dp[i] = 0
					}
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
