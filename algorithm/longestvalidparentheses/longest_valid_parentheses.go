package longestvalidparentheses

func longestValidParentheses(s string) int {
	stack := make([]int, 0)

	for i := 0; i < len(s); i++ {
		size := len(stack)
		if size >= 1 {
			top := stack[len(stack)-1]
			if (s[i] == ')' && s[top] == '(') || (s[i] == '}' && s[top] == '{') || (s[i] == ']' && s[top] == '[') {
				//出栈
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}
		} else {
			stack = append(stack, i)
		}
	}

	stack = append([]int{-1}, stack...)
	stack = append(stack, len(s))

	longest := 0
	for i := 1; i < len(stack); i++ {
		gap := stack[i] - stack[i-1] - 1
		if longest < gap {
			longest = gap
		}
	}
	return longest
}
