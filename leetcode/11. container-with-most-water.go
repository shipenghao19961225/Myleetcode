package leetcode

// 从两边向中间延展，若左边高度小，挪右边的是没用的，此时挪左边，右边
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		if height[left] < height[right] {
			res = max(res, height[left]*(right-left))
			left++
		} else {
			res = max(res, height[right]*(right-left))
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
