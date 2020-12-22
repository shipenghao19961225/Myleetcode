package leetcode

// 第一种方法为动态规划，时间复杂度为O(n^2)
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	maxLen := 1
	start := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if i-j < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i-1][j+1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] == true && i-j+1 > maxLen {
				maxLen = i - j + 1
				start = j
			}
		}
	}
	return s[start : start+maxLen]
}

// 第二种方法，为中心扩散法，利用多返回值的特性，可以写出简练的代码
// 时间复杂度为O(n^2)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		tmpStart, tmpEnd := expand(s, i, i)
		if tmpEnd-tmpStart > end-start {
			start, end = tmpStart, tmpEnd
		}
		tmpStart, tmpEnd = expand(s, i, i+1)
		if tmpEnd-tmpStart > end-start {
			start, end = tmpStart, tmpEnd
		}
	}
	return s[start : end+1]
}

func expand(s string, start, end int) (int, int) {
	for ; start >= 0 && end < len(s) && s[start] == s[end]; start, end = start-1, end+1 {
	}
	return start + 1, end - 1
}
