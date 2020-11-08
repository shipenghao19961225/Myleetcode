package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// 这里可以提前终止循环，降低时间
		if nums[i] > 0 {
			break
		}
		// 总是忘记，第一项也有可能是重复项，所以需要判断
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > 0 {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return res
}
