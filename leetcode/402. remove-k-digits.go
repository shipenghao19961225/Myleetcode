package leetcode

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && num[i] < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		if len(stack) == 0 && num[i] == '0' {
			continue
		}
		stack = append(stack, num[i])
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
