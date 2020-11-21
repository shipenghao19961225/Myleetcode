package leetcode

func maxSubArray(nums []int) int {
	dp := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if dp < 0 {
			dp = nums[i]
		} else {
			dp = dp + nums[i]
		}
		res = max(res, dp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
