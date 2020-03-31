package threeSumClosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	diff := int(^uint(0) >> 1)
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		low, high := i+1, len(nums)-1
		for low < high {
			cur := nums[i] + nums[low] + nums[high]
			if cur > target {
				high--
				if cur-target < diff {
					diff = cur - target
					result = cur
				}
			} else if cur < target {
				low++
				if target-cur < diff {
					diff = target - cur
					result = cur
				}
			} else {
				return target
			}
		}
	}

	return result
}
