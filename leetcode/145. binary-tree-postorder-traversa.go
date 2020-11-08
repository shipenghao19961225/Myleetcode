package leetcode

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	m := make(map[*TreeNode]bool)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		for len(stack) > 0 && m[stack[len(stack)-1]] {
			res = append(res, stack[len(stack)-1].Val)
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1].Right
			m[stack[len(stack)-1]] = true
		}
	}
	return res
}

// 倾向于第二种解法，更容易记住
func postorderTraversal2(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{topNode.Val}, res...)
		if topNode.Left != nil {
			stack = append(stack, topNode.Left)
		}
		if topNode.Right != nil {
			stack = append(stack, topNode.Right)
		}
	}
	return res
}
