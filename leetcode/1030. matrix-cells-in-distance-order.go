package leetcode

// 方法一 bfs
type point struct {
	x int
	y int
}

var dirX = []int{0, 1, 0, -1}
var dirY = []int{1, 0, -1, 0}
var visited map[point]bool

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	visited = make(map[point]bool)
	length := R * C
	queue := make([]*point, 0, length)
	res := make([][]int, 0, length)
	queue = append(queue, &point{r0, c0})
	visited[point{r0, c0}] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			topPoint := queue[0]
			queue = queue[1:]
			res = append(res, []int{topPoint.x, topPoint.y})
			for i := 0; i < 4; i++ {
				newX := topPoint.x + dirX[i]
				newY := topPoint.y + dirY[i]
				if inBound(newX, newY, R, C) && !visited[point{newX, newY}] {
					visited[point{newX, newY}] = true
					queue = append(queue, &point{newX, newY})
				}
			}
		}
	}
	return res
}

func inBound(x, y, R, C int) bool {
	return x >= 0 && x < R && y >= 0 && y < C
}

// 方法二 桶排序
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0, R*C)
	bucket := make(map[int][][]int)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			d := calDistance(i, j, r0, c0)
			bucket[d] = append(bucket[d], []int{i, j})
		}
	}
	for i := 0; i < R+C; i++ {
		for _, point := range bucket[i] {
			res = append(res, point)
		}
	}
	return res
}
func calDistance(i, j, r0, c0 int) int {
	return abs(i-r0) + abs(j-c0)
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
