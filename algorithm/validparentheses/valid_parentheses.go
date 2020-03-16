package validparentheses

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		size := len(stack)
		if size >= 1 {
			top := stack[len(stack)-1]
			if (s[i] == ')' || s[i] == ']' || s[i] == '}') && (top == ')' || top == ']' || top == '}') {
				return false
			}
			if (s[i] == ')' && top == '(') || (s[i] == '}' && top == '{') || (s[i] == ']' && top == '[') {
				//出栈
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isValidV2(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		size := len(stack)
		if size > 1 {
			top := stack[size-2]
			if top == ']' || top == ')' || top == '}' {
				return false
			}
			if s[i] == ']' && stack[size-2] == '[' {
				stack = stack[:size-2]
			}
			if s[i] == ')' && stack[size-2] == '(' {
				stack = stack[:size-2]
			}
			if s[i] == '}' && stack[size-2] == '{' {
				stack = stack[:size-2]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
