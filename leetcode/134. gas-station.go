package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	curSum, tank, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		tank += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if tank >= 0 {
		return start
	}
	return -1
}

// @Charles000 假设把无穷多个加油站排在一起。 规定只能往右。从0开始计算这个无穷的序列上每个点的剩余油量，剩余量允许为负数。
// 有解的情况等价于， 从某个位置i开始， 连续n个位置的剩余量不小于位置i的剩余量。

// 会有三种情况：

// 走完一遍以后，油量增加
// 走完一遍以后油量不变
// 走完一遍以后油量减少
// 1， 一定能找到解
// 2，这种情况下， 解存在，即 “周期函数一定有最大值和最小值。 只要从最小值出发，永远不会低于初值。”
// 3， 一定无解
