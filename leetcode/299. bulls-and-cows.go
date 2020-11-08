package leetcode

import "strconv"

func getHint(secret string, guess string) string {
	m := make([]int, 10)
	x, y := 0, 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			x++
		} else {
			b1 := secret[i] - '0'
			b2 := guess[i] - '0'
			if m[b1] < 0 {
				y++
			}
			m[b1]++
			if m[b2] > 0 {
				y++
			}
			m[b2]--
		}
	}
	return strconv.Itoa(x) + "A" + strconv.Itoa(y) + "B"
}
