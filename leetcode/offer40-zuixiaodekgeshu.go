package leetcode

import "container/heap"

// 第一种方法，大根堆。
// 这里实现的是go container中的heap接口，直接使用go container api, 注意
// 这里求的是最小的K个数，所以需要构建大小为K的大根堆，当元素小于堆顶元素时
// 堆顶弹出，新元素压入，heap并没有提供访问堆顶元素的API,经检验，(*h)[0]就为
// 堆顶元素。时间复杂度为nlogk,比较慢。
type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	res := make([]int, 0, k)
	tmp := make([]int, 0, k)
	heapTmp := myHeap(tmp)
	heap1 := &heapTmp
	for i := 0; i < k; i++ {
		heap.Push(heap1, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if (*heap1)[0] > arr[i] {
			heap.Pop(heap1)
			heap.Push(heap1, arr[i])
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, (*heap1)[i])
	}
	return res
}

func getLeastNumbers1(arr []int, k int) []int {
	// 这里的鲁棒性，因为k=0会一直逼近到-1，导致索引越界
	if k == 0 {
		return []int{}
	}
	return helper(arr, 0, len(arr)-1, k)
}

func helper(a []int, lo, hi int, k int) []int {
	index := partition(a, lo, hi)
	if index+1 == k {
		return a[:index+1]
	} else if index+1 < k {
		return helper(a, index+1, hi, k)
	} else {
		return helper(a, lo, index-1, k)
	}
}

// 第二种方法
// 这种方法其实就是利用了快排中的partition思想，将数组一分为2
// 选取一个元素作为pivot,在该元素之前的所有元素都比其小

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	return helper(arr, 0, len(arr)-1, k)
}
func helper1(arr []int, lo, hi, k int) []int {
	index := partition(arr, lo, hi)
	if index+1 == k {
		return arr[:k]
	} else if index+1 < k {
		return helper1(arr, index+1, hi, k)
	} else {
		return helper1(arr, lo, index-1, k)
	}
}
func partition(arr []int, lo, hi int) int {
	j := lo
	pivot := arr[hi]
	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	arr[j], arr[hi] = arr[hi], arr[j]
	return j
}
