package knapsack

import "fmt"

//背包问题：动态规划，不涉及商品的价值。

type IBag interface {
	Solve(items []int, n, w int) int
}

//用二维数组states [][]bool来存放状态
type Bag struct{}

func (b *Bag) Solve(items []int, n, w int) int {
	//初始化状态转移表
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, w+1)
	}

	//初始化第0行的状态
	states[0][0] = false
	if items[0] <= w {
		states[0][items[0]] = true
	}

	//开始状态转移，从第一行开始
	for i := 1; i < n; i++ {
		//不选第i个
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		//选第i个
		for j := 0; j <= w-items[i]; j++ {
			if states[i-1][j] {
				states[i][j+items[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}
	fmt.Println(states)
	return 0
}

//用一维数组states []bool来存放状态
type BagV1 struct{}

func (b *BagV1) Solve(items []int, n, w int) int {
	states := make([]bool, w+1)
	states[0] = true
	if items[0] <= w {
		states[items[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := w - items[i]; j >= 0; j-- {
			if states[j] {
				states[j+items[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}
