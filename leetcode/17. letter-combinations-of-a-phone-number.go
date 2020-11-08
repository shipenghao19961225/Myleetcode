package leetcode

var m map[byte]string
var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) == 0 {
		return res
	}

	m = make(map[byte]string)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"
	dfs("", digits)
	return res
}
func dfs(path, digits string) {
	// termination
	if len(path) == len(digits) {
		res = append(res, path)
		return
	}
	//process
	choices := m[digits[len(path)]]
	for i := 0; i < len(choices); i++ {
		dfs(path+string(choices[i]), digits)
	}
}
