package leetcode

func exist(board [][]byte, word string) bool {
	res := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res = dfs(i, j, 0, word, board)
			if res == true {
				break
			}
		}
		if res == true {
			break
		}
	}
	return res
}
func dfs(i, j, idx int, word string, board [][]byte) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	// process
	if board[i][j] != word[idx] {
		return false
	}
	board[i][j] = '#'
	if dfs(i+1, j, idx+1, word, board) ||
		dfs(i-1, j, idx+1, word, board) ||
		dfs(i, j+1, idx+1, word, board) ||
		dfs(i, j-1, idx+1, word, board) {
		return true
	}
	board[i][j] = word[idx]
	return false
}
