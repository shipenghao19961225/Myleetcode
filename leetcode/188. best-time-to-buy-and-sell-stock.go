package leetcode

// 买卖股票问题四，允许进行k次买卖
// T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
// T[i][k][1] = max(T[i-1][k][1], T[i-1][k-1][0] - prices[i])
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	if k >= n/2 {
		return maxProfit2(prices)
	}
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 1; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}
	// 注意这里的压缩空间方法，比较讲究，倒过来遍历
	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}
	return dp[k][0]

}
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	s0, s1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		tmp := s0
		s0 = max(s0, s1+prices[i])
		s1 = max(s1, tmp-prices[i])
	}
	return s0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
