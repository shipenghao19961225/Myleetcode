package leetcode

func isLongPressedName(name string, typed string) bool {
	j := 0
	for i := 0; i < len(typed); i++ {
		if j < len(name) && typed[i] == name[j] {
			j++
		} else if i == 0 || typed[i] != typed[i-1] {
			return false
		}
	}
	if j == len(name) {
		return true
	}
	return false

}
