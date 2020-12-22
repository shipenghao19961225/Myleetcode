package leetcode

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	count := 0
	for i >= 0 && s[i] != ' ' {
		i--
		count++
	}
	return count
}
