package leetcode

// https://leetcode-cn.com/problems/house-robber/solution/da-jia-jie-she-dong-tai-gui-hua-jie-gou-hua-si-lu-/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, cur := 0, 0
	for i := 0; i < len(nums); i++ {
		tmp := cur
		cur = max(pre+nums[i], cur)
		pre = tmp
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 另一种解法，覃超老师说这种二维三维的动态规划往往更加常用
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
