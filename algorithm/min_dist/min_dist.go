package min_dist

import "math"

type IMinDist interface {
	MinDistDP(matrix [][]int, n int) int
}

//状态转仪表
type DP struct{}

//记录当前最短距离
var minDist = math.MaxInt64

func (d *DP) MinDistDP(matrix [][]int, n int) int {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, n)
	}
	states[0][0] = matrix[0][0]
	//初始化第一列
	for i := 1; i < n; i++ {
		states[i][0] = states[i-1][0] + matrix[i][0]
	}
	//初始化第一行
	for j := 1; j < n; j++ {
		states[0][j] = states[0][j-1] + matrix[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if states[i-1][j] > states[i][j-1] {
				states[i][j] = states[i][j-1] + matrix[i][j]
			} else {
				states[i][j] = states[i-1][j] + matrix[i][j]
			}
		}
	}

	return states[n-1][n-1]
}

//状态转移表
func MinDistDP(i, j, n int, matrix [][]int) int {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, n)
	}
	if i == 0 && j == 0 {
		return matrix[0][0]
	}
	if states[i][j] > 0 {
		return states[i][j]
	}

	minLeft := math.MaxInt32
	if j >= 1 {
		minLeft = MinDistDP(i, j-1, n, matrix)
	}
	minUp := math.MaxInt32
	if i >= 1 {
		minUp = MinDistDP(i-1, j, n, matrix)
	}
	currMinDist := int(math.Min(float64(minUp), float64(minLeft))) + matrix[i][j]
	states[i][j] = currMinDist
	return currMinDist
}
