package leetcode

func invertTree(root *TreeNode) *TreeNode {
	// termination
	if root == nil {
		return nil
	}
	// process current level
	root.Left, root.Right = root.Right, root.Left
	// drill down
	invertTree(root.Left)
	invertTree(root.Right)
	return root
	// reverse this level state if needed
}
