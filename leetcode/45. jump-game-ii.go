package leetcode

func jump(nums []int) int {
	step := 0
	maxReach := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == end {
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
