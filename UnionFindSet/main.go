package main

func main() {

}

type UnionSet struct {
	Father []int
}

func (u UnionSet) Init() {
	for i := range u.Father {
		u.Father[i] = i
	}
}
func (u UnionSet) find(x int) int {
	if x == u.Father[x] {
		return x
	}
	//路径压缩
	u.Father[x] = u.find(u.Father[x])
	return u.Father[x]
}

func (u UnionSet) isSame(x int, y int) bool {
	return u.find(x) == u.find(y)
}

func (u UnionSet) join(x int, y int) {
	x = u.find(x)
	y = u.find(y)
	if x == y {
		return
	}
	u.Father[y] = x
}

func isTreeAfterRemoveEdge(u UnionSet, edges [][]int, deleteEdge int) bool {
	for i, edge := range edges {
		if i == deleteEdge {
			continue
		}
		if u.isSame(edge[0], edge[1]) {
			return false
		}
		u.join(edge[0], edge[1])
	}
	return true
}
func getRemoveEdge(u UnionSet, edges [][]int) []int {
	for _, edge := range edges {
		if u.isSame(edge[0], edge[1]) {
			return edge
		}
		u.join(edge[0], edge[1])
	}
	return nil
}
func findRedundantDirectedConnection(edges [][]int) []int {
	inDegree := make([]int, 1001)
	for i := 0; i < len(edges); i++ {
		inDegree[edges[i][1]]++
	}
	vec := make([]int, 0)
	for i := len(edges) - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			vec = append(vec, i)
		}
	}
	u := UnionSet{
		Father: make([]int, 1001),
	}
	u.Init()
	if len(vec) > 0 {
		if isTreeAfterRemoveEdge(u, edges, vec[0]) {
			return edges[vec[0]]
		} else {
			return edges[vec[1]]
		}
	}
	u.Init()
	return getRemoveEdge(u, edges)
}
func findRedundantConnection(edges [][]int) []int {
	u := UnionSet{
		Father: make([]int, 1001),
	}
	u.Init()
	for _, edge := range edges {
		if u.isSame(edge[0], edge[1]) {
			return edge
		} else {
			u.join(edge[0], edge[1])
		}
	}
	return nil
}
func validPath(n int, edges [][]int, source int, destination int) bool {
	u := UnionSet{
		Father: make([]int, n),
	}
	u.Init()
	for _, edge := range edges {
		u.join(edge[0], edge[1])
	}
	return u.isSame(source, destination)
}
