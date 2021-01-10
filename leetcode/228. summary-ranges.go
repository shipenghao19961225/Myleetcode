package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	start := 0
	res := []string{}
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if nums[start] == nums[i] {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i]))
			}
			start = i + 1
		}
	}
	return res
}
