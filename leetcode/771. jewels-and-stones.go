package leetcode

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[rune]bool)
	for _, char := range jewels {
		m[char] = true
	}
	count := 0
	for _, char := range stones {
		if _, ok := m[char]; ok {
			count++
		}
	}
	return count
}
