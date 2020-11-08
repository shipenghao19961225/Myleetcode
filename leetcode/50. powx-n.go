package leetcode

// 分治解法
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / Pow(x, -n)
	}
	return Pow(x, n)
}
func Pow(x float64, n int) float64 {
	// termination
	if n == 0 {
		return 1
	}
	// drill down
	tmp := Pow(x, n/2)
	// process
	if n%2 == 1 {
		return tmp * tmp * x
	}
	return tmp * tmp
}

// 快速幂解法
func myPow(x float64, n int) float64 {
	flag := 0
	cont := x
	var res float64 = 1
	if n < 0 {
		n = -1 * n
		flag = 1
	}
	for n > 0 {
		if n&1 == 1 {
			res *= cont
		}
		cont = cont * cont
		n = n >> 1
	}
	if flag == 1 {
		return 1 / res
	}
	return res
}
