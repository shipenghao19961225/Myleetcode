package thoughts

func mergeSort(a []int) []int {
	mergeSortSub(a, 0, len(a)-1)
	return a
}
func mergeSortSub(a []int, start, end int) {
	// termination
	if start >= end {
		return
	}
	// divide and drill down
	mid := (start + end) / 2
	mergeSortSub(a, start, mid)
	mergeSortSub(a, mid+1, end)
	// merge the result
	merge(a[start:mid+1], a[mid+1:end+1], a[start:end+1])
	// restore
}
func merge(a1, a2, a3 []int) {
	tmp := make([]int, len(a3))
	i, j, k := 0, 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			tmp[k] = a1[i]
			k++
			i++
		} else {
			tmp[k] = a2[j]
			k++
			j++
		}
	}
	for i < len(a1) {
		tmp[k] = a1[i]
		k++
		i++
	}
	for j < len(a2) {
		tmp[k] = a2[j]
		k++
		j++
	}
	for m := 0; m < len(a3); m++ {
		a3[m] = tmp[m]
	}
}
