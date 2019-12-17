package threeSum

import (
	"sort"
	"strconv"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 找出所有满足条件且不重复的三元组。
//
//来源：LeetCode
//链接：https://leetcode-cn.com/problems/3sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([]int, 3)
	total := make([][]int, 0)
	solveThreeSum(0, 0, 0, len(nums), nums, result, &total)
	return deduplicate(total)
}

//考察第K个，used表示当前选了多少个。当前和为sum, m表示一共有多少个数，result为存储了目前为止的部分解, total表示所有解
func solveThreeSum(k, used, sum, m int, nums, result []int, total *[][]int) {
	if used == 3 {
		if sum == 0 {
			tmp := make([]int, len(result))
			copy(tmp, result)
			*total = append(*total, tmp)
		}
		return
	}
	if k == m {
		return
	}
	if sum > 0 && nums[k] > 0 {
		return
	}

	//不选第k个
	solveThreeSum(k+1, used, sum, m, nums, result, total)
	//选第k个
	result[used] = nums[k]
	used++
	solveThreeSum(k+1, used, sum+nums[k], m, nums, result, total)
}

func deduplicate(total [][]int) [][]int {
	result := make([][]int, 0)
	tmp := make(map[string][]int)
	for _, v := range total {
		sort.Ints(v)
		vStr := ""
		for _, num := range v {
			vStr += strconv.Itoa(num)
		}
		tmp[vStr] = v
	}
	for _, v := range tmp {
		result = append(result, v)
	}
	return result
}
