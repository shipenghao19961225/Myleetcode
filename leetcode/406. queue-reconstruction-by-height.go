package leetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i])
		people[p[1]] = p
	}
	return people
}
