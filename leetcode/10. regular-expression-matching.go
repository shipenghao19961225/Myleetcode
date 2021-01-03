package leetcode

// https://leetcode-cn.com/problems/regular-expression-matching/solution/dong-tai-gui-hua-zui-xiang-xi-jie-da-you-jian-zhi-
// 这道题还是蛮难的。
//
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = p[i-1] == '*' && dp[i-2][0]
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			nowS, nowP := s[j-1], p[i-1]
			if nowS == nowP || nowP == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if nowP == '*' && i >= 2 {
				lastP := p[i-2]
				if lastP == nowS || lastP == '.' {
					dp[i][j] = dp[i-1][j] || dp[i][j-1]
				}
				dp[i][j] = dp[i][j] || dp[i-2][j]
			}
		}
	}
	return dp[m][n]
}
