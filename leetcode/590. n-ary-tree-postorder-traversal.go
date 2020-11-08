package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	stack := []*Node{}
	res := []int{}
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, topNode.Val)
		for i := 0; i < len(topNode.Children); i++ {
			stack = append(stack, topNode.Children[i])
		}
	}
	reverse(res)
	return res
}

// 在数组之前插入操作时间复杂度低，这里直接用反转，降低时间复杂度
func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
