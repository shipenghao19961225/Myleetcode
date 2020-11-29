package leetcode

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			i--
			p2--
		}
	}
}
