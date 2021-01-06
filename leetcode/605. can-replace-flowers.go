package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i++
		} else if flowerbed[i] == 1 {
			i++
		}
	}
	return n <= 0
}
