package leetcode

// 这里并查集的Find路径压缩都压缩到根节点的子结点，还有一种是对半压缩的，见200.岛屿数量
// 其实两者压缩都可以，但是这个题，似乎只有第一种压缩解法。难度蛮大的，不过并查集理解的更深了点

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	u := new(UnionFind)
	u.Init(2 * n)
	m := make(map[string]int)

	// 用来标识字母
	id := 0

	// 将数字的关系添加到并查集中
	for i := 0; i < n; i++ {
		if _, ok := m[equations[i][0]]; !ok {
			m[equations[i][0]] = id
			id++
		}
		if _, ok := m[equations[i][1]]; !ok {
			m[equations[i][1]] = id
			id++
		}
		u.Union(m[equations[i][0]], m[equations[i][1]], values[i])
	}

	// 查询每个数字的关系
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		s0, s1 := queries[i][0], queries[i][1]
		id0, ok0 := m[s0]
		id1, ok1 := m[s1]
		if !ok0 || !ok1 {
			res[i] = -1.0
		} else {
			res[i] = u.isConnected(id0, id1)
		}
	}
	return res
}

type UnionFind struct {
	parent []int

	// 这个字段是标识到父节点的权重
	weight []float64
}

func (u *UnionFind) Init(n int) {
	u.parent = make([]int, n)
	u.weight = make([]float64, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.weight[i] = 1.0
	}
}

func (u *UnionFind) Union(x, y int, value float64) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return
	}
	u.parent[rootX] = rootY
	// 这个想一下平行四边形
	u.weight[rootX] = u.weight[y] * value / u.weight[x]
}

func (u *UnionFind) Find(x int) int {
	if u.parent[x] != x {
		origin := u.parent[x]
		//将所有经过的结点都连接到父节点上
		u.parent[x] = u.Find(u.parent[x])
		// 更新权重，父结点的权重，乘自己的weight
		u.weight[x] *= u.weight[origin]
	}
	return u.parent[x]
}

func (u *UnionFind) isConnected(x, y int) float64 {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return u.weight[x] / u.weight[y]
	}
	return -1.0
}
