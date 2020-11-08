package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	// 维护一个双端单调递减队列，首元素是最大的

	deque := make([]int, 0, len(nums))
	for i, num := range nums {
		if len(deque) > 0 && i-k >= deque[0] {
			deque = deque[1:]
		}
		for len(deque) > 0 && num > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i-k+1 >= 0 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
