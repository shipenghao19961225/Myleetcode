package leetcode

func removeOuterParentheses1(S string) string {
	stack := make([]byte, 0, len(S))
	res := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == '(' && len(stack) == 0 {
			stack = append(stack, '(')
		} else if S[i] == '(' && len(stack) != 0 {
			stack = append(stack, '(')
			res = append(res, '(')
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				res = append(res, ')')
			}
		}
	}
	return string(res)
}

// 第二种解法，很妙！！~~~~~
func removeOuterParentheses2(S string) string {
	res := make([]byte, 0, len(S))
	level := 0
	for i := 0; i < len(S); i++ {
		if S[i] == ')' {
			level--
		}
		if level >= 1 {
			res = append(res, S[i])
		}
		if S[i] == '(' {
			level++
		}
	}
	return string(res)
}
