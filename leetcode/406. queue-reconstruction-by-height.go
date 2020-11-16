package leetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0, len(people))
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, person := range people {
		index := person[1]
		res = append(res[:index], append([][]int{person}, res[index:]...)...)
	}
	return res
}
