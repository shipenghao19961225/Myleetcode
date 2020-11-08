package leetcode

// 这道题三种方法，方法一，排序直接找n/2元素
// 	方法二，快排partition找
// 方法三， 摩尔投票

// 方法二
func majorityElement(nums []int) int {
	return helper(nums, 0, len(nums)-1, len(nums)/2)
}

func helper(nums []int, start, end, target int) int {
	index := partition(nums, start, end)
	if index == target {
		return nums[target]
	}
	if index > target {
		return helper(nums, start, index-1, target)
	}
	return helper(nums, index+1, end, target)

}
func partition(nums []int, start, end int) int {
	pivot := nums[end]
	j := start
	for i := start; i < end; i++ {
		if nums[i] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[end] = nums[end], nums[j]
	return j
}

// 方法三
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	major, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count++
		} else if nums[i] == major {
			count++
		} else {
			count--
		}
	}
	return major
}
