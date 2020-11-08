package leetcode

import "math"

// 该题可以用中序遍历解，也可以用递归解

// 中序遍历解法
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	cur := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if topNode.Val <= cur {
			return false
		}
		root = topNode.Right
		cur = topNode.Val
	}
	return true

}

// 递归解法
func isValidBST2(root *TreeNode) bool {
	return dfs2(root, math.MinInt64, math.MaxInt64)
}

func dfs2(root *TreeNode, low, high int) bool {
	// termination
	if root == nil {
		return true
	}
	// process
	if root.Val >= high || root.Val <= low {
		return false
	}
	// drill down
	return dfs2(root.Left, low, root.Val) && dfs2(root.Right, root.Val, high)
}
