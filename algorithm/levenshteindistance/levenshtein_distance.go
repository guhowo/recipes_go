package levenshteindistance

import (
	"math"
)

func LevenshteinDistance(a, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return int(math.Max(float64(lenA), float64(lenB)))
	}

	lev := make([][]int, lenA+1)
	for i := 0; i <= lenA; i++ {
		lev[i] = make([]int, lenB+1)
	}

	//初始化第一列
	for i := 0; i <= lenA; i++ {
		lev[i][0] = i
	}
	//初始化第一行
	for j := 0; j <= lenB; j++ {
		lev[0][j] = j
	}

	//计算
	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			if a[i-1] == b[j-1] {
				lev[i][j] = lev[i-1][j-1]
			} else {
				lev[i][j] = int(math.Min(math.Min(float64(lev[i-1][j]), float64(lev[i][j-1])), float64(lev[i-1][j-1]))) + 1
			}
		}
	}

	return lev[lenA][lenB]
}
