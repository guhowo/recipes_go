package wildcard_matching

//状态 dp[i][j] : 表示 s 的前 i 个字符和 p 的前 j 个字符是否匹配 (true 的话表示匹配)
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	//初始化第一行
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] && (p[i-1] == '*')
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
			//第i,j为不匹配，那么改行后面的都不匹配
			if dp[i][j] == false {
				continue
			}
		}
	}

	return dp[m][n]
}
