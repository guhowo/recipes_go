package levenshteindistance

import "math"

func LevenshteinDistance(a, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return int(math.Max(float64(lenA), float64(lenB)))
	}

	lev := make([][]int, lenA)
	for i := 0; i < lenA; i++ {
		lev[i] = make([]int, lenB)
	}

	//初始化顶点[0][0]
	if a[0] == b[0] {
		lev[0][0] = 0
	} else {
		lev[0][0] = 1
	}
	//初始化第一列
	for i := 1; i < lenA; i++ {
		if a[i] == b[0] {
			lev[i][0] = lev[i-1][0]
		} else {
			lev[i][0] = lev[i-1][0] + 1
		}
	}
	//初始化第一行
	for j := 1; j < lenB; j++ {
		if a[0] == b[j] {
			lev[0][j] = lev[0][j-1]
		} else {
			lev[0][j] = lev[0][j-1] + 1
		}
	}

	//计算
	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			if a[i] == b[j] {
				lev[i][j] = lev[i-1][j-1]
			} else {
				lev[i][j] = int(math.Min(math.Min(float64(lev[i-1][j]), float64(lev[i][j-1])), float64(lev[i-1][j-1]))) + 1
			}
		}
	}

	return lev[lenA-1][lenB-1]
}
