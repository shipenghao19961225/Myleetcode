package leetcode

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	// 增加哨兵可以减少边界条件的处理，很常用
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	stack := make([]int, n+2)
	res := 0
	for i, height := range heights {
		for height < heights[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = max(res, heights[topIndex]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return res
}
