package rmatch

var matched = false

func rmatch(ti, pj int, text, pattern string) {
	if matched {
		return
	}

	//到达pattern末尾，如果text也到末尾，表示匹配，否则不匹配
	if pj == len(pattern) {
		if ti == len(text) {
			matched = true
		}
		return
	}

	//如果是*，匹配0个或任意多个字符。代码从0长度开始匹配
	if pattern[pj] == '*' {
		for k := 0; k <= len(text)-ti; k++ {
			rmatch(ti+k, pj+1, text, pattern)
		}
	} else if pattern[pj] == '?' {
		//如果是？，匹配0个或1个字符。
		//匹配0个字符
		rmatch(ti, pj+1, text, pattern)
		//匹配1个字符
		rmatch(ti+1, pj+1, text, pattern)
	} else if pattern[pj] == '.' {
		rmatch(ti+1, pj+1, text, pattern)
	} else if ti < len(text) && pattern[pj] == text[ti] {
		//纯字符串匹配
		rmatch(ti+1, pj+1, text, pattern)
	}

}
