package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for _, v1 := range nums1 {
		m[v1]++
	}
	for _, v2 := range nums2 {
		if value, ok := m[v2]; ok && value > 0 {
			m[v2]--
			res = append(res, v2)
		}
	}
	return res
}
