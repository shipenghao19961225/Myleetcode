package leetcode

var res [][]int

func subsets(nums []int) [][]int {
	res = [][]int{}
	dfs(0, []int{}, nums)
	return res

}

func dfs(index int, path, nums []int) {
	// termination
	if index == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	// process drill down
	dfs(index+1, path, nums)
	path = append(path, nums[index])
	dfs(index+1, path, nums)
	// reverse
	path = path[:len(path)-1]
}
