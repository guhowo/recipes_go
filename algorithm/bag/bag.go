package bag

var maxW int //当前求得的最大的总重量
var n int    //items的个数
var result []int

//考察第k个，items表示每个物品的重量，总共有n个物品，背包可承受的总重量为w,cw表示当前总重量（不含k）
func getMaxW(k, cw, w int, items []int) {
	if k == n || cw == w {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	getMaxW(k+1, cw, w, items) //第k个不选，考察第k+1个
	if cw+items[k] <= w {
		getMaxW(k+1, cw+items[k], w, items)
		result[k] = 1
	}
}
