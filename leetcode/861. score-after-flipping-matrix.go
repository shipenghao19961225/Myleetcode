package leetcode

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	res := m * 1 << (n - 1)
	for j := 1; j < n; j++ {
		nOnes := 0
		for i := 0; i < m; i++ {
			if A[i][0] == A[i][j] {
				nOnes++
			}
		}
		k := max(nOnes, m-nOnes)
		res += k * 1 << (n - j - 1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
