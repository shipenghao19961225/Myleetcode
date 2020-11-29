package leetcode

// 1. 回溯加记忆化
var cache [][]int
var visited [][]bool

func longestCommonSubsequence(text1 string, text2 string) int {
	cache = make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		cache[i] = make([]int, len(text2))
	}
	visited = make([][]bool, len(text1))
	for i := 0; i < len(text1); i++ {
		visited[i] = make([]bool, len(text2))
	}
	return dfs(0, 0, text1, text2)
}
func dfs(i, j int, text1, text2 string) int {
	// termination
	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if visited[i][j] {
		return cache[i][j]
	}
	// process and drill
	if text1[i] == text2[j] {
		cache[i][j] = dfs(i+1, j+1, text1, text2) + 1
		visited[i][j] = true
		return cache[i][j]
	}
	cache[i][j] = max(dfs(i+1, j, text1, text2), dfs(i, j+1, text1, text2))
	visited[i][j] = true
	return cache[i][j]
	// reverse the state
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 2. dp
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text1)+1)
	m := len(text2) + 1
	n := len(text1) + 1

	for i := 1; i < m; i++ {
		follow := dp[0]
		for j := 1; j < n; j++ {
			tmp := follow
			follow = dp[j]
			if text1[j-1] == text2[i-1] {
				dp[j] = tmp + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
