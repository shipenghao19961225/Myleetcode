package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 1 {
		return nums[0]
	}
	nums1 := nums[1:]
	pre, cur := 0, 0
	for i := 0; i < len(nums1); i++ {
		tmp := cur
		cur = max(pre+nums1[i], cur)
		pre = tmp
	}
	nums2 := nums[:len(nums)-1]
	pre, cur2 := 0, 0
	for i := 0; i < len(nums2); i++ {
		tmp := cur2
		cur2 = max(pre+nums2[i], cur2)
		pre = tmp
	}
	return max(cur, cur2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
