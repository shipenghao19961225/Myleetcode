package leetcode

func fib(n int) int {
	if n <= 1 {
		return n
	}
	pre, cur := 0, 1
	for i := 2; i <= n; i++ {
		newcur := pre + cur
		newpre := cur
		pre = newpre
		cur = newcur
	}
	return cur
}
