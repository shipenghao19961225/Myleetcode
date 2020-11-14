package leetcode

// 牛顿迭代法
func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

// 二分法
func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	low, high := 0, x+1
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
