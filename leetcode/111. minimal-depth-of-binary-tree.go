package leetcode

import "math"

// 这道题，可以深度优先搜索做，也可以广度优先搜索做
// 深度优先搜索做的话，需要进行剪枝操作
// 广度优先搜索，就是层次遍历，自动进行了一个剪枝

// 已记录最小深度

var minDe int

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDe = math.MaxInt32
	dfs1(root, 1)
	return minDe
}

func dfs1(root *TreeNode, depth int) {
	// termination

	if root.Left == nil && root.Right == nil {
		// 如果到了子结点，就取 当前深度 和 已经记录深度的最小值 存储为 已记录深度的最小值
		minDe = min(depth, minDe)
	}
	// process
	// down
	// 这里进行个剪枝操作，如果递归时当前深度小于已经记录的最小深度，才可以继续走下去
	// 如果当前深度已经大于所记录的最小深度，就没必要走下去
	if depth < minDe {
		if root.Left != nil {
			dfs1(root.Left, depth+1)
		}
		if root.Right != nil {
			dfs1(root.Right, depth+1)
		}
	}
}

// 广度优先搜索，类似于层次遍历
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i].Left == nil && q[i].Right == nil {
				return level
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[size:]
		level++
	}
	return level
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
