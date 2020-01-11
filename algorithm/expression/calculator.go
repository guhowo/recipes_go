package expression

import (
	"strconv"
)

func Calculate(expression string) int {
	nums := make([]int, 0)
	operators := make([]string, 0)
	keys := split(expression)
	for _, key := range keys {
		if key == "*" || key == "/" || key == "+" || key == "-" {
			for isLower(key, operators) {
				op := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				before, after := nums[len(nums)-2], nums[len(nums)-1]
				nums = nums[:len(nums)-2]
				if op == "+" {
					nums = append(nums, before+after)
				}
				if op == "-" {
					nums = append(nums, before-after)
				}
				if op == "*" {
					nums = append(nums, before*after)
				}
				if op == "/" {
					nums = append(nums, before/after)
				}
			}
			operators = append(operators, key)
		} else {
			num, _ := strconv.Atoi(key)
			nums = append(nums, num)
		}
	}

	//最后一次运算需要做
	before, after := nums[0], nums[1]
	nums = nums[:len(nums)-2]
	if operators[0] == "+" {
		nums = append(nums, before+after)
	}
	if operators[0] == "-" {
		nums = append(nums, before-after)
	}
	if operators[0] == "*" {
		nums = append(nums, before*after)
	}
	if operators[0] == "/" {
		nums = append(nums, before/after)
	}
	return nums[0]
}

func split(s string) []string {
	result := make([]string, 0)
	var i, j int
	for j = -1; i < len(s); i++ {
		if s[i] != '*' && s[i] != '/' && s[i] != '+' && s[i] != '-' {
			continue
		} else {
			result = append(result, s[j+1:i])
			result = append(result, string(s[i]))
			j = i
		}
	}
	result = append(result, s[j+1:i])
	return result
}

//待压入的运算符优先级是否比op栈顶的优先级低或者相等。是返回true
func isLower(key string, op []string) bool {
	if len(op) != 0 {
		if key == "+" || key == "-" {
			return true
		}
		if (key == "*" || key == "/") && (op[len(op)-1] == "*" || op[len(op)-1] == "/") {
			return true
		}
	}
	return false
}
