package leetcode

// 牛顿迭代法
func isPerfectSquare(num int) bool {
	x := num
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}

// 二分法
func isPerfectSquare(num int) bool {
	low, high := 0, num
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
