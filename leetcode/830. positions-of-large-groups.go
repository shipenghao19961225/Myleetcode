package leetcode

func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0, len(s)/3+1)
	count := 1
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || s[i+1] != s[i] {
			if count >= 3 {
				res = append(res, []int{i - count + 1, i})
			}
			count = 1
		} else {
			count++
		}
	}
	return res
}
