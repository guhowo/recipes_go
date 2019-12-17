package knapsack3

type IBag interface {
	Solve(weight []int, values []int, n, w int) int
}

type Bag struct{}

/*
 weight：物品重量
 values： 物品价值
 n：物品个数
 w：背包承重最大值
*/
func (b *Bag) Solve(weight []int, values []int, n, w int) int {
	// states保存考察第i个物品时，放和不放，能达到的价值
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, w+1)
	}

	//初始化状态表为-1
	for i := 0; i < n; i++ {
		for j := 0; j < w+1; j++ {
			states[i][j] = -1
		}
	}

	//初始化第0行：第0个选和不选
	states[0][0] = 0
	if weight[0] <= w {
		states[0][weight[0]] = values[0]
	}

	//求解第1~n-1个物品
	for i := 1; i < n; i++ {
		//不选第i个物品
		for j := 0; j < w+1; j++ {
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}

		//选第i个物品
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + values[i]
				if v > states[i][j+weight[i]] {
					states[i][j+weight[i]] = v
				}
			}
		}
	}

	maxValue := -1
	for j := 0; j < w+1; j++ {
		if maxValue < states[n-1][j] {
			maxValue = states[n-1][j]
		}
	}
	return maxValue
}
