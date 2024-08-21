package main

import (
	"fmt"
	"sort"
)

var father []int

type edge struct {
	V1, V2, Val int
}

func main() {
	var v, e int
	fmt.Scanf("%d %d\n", &v, &e)
	father = make([]int, v+1)
	edges := make([]edge, e)
	// 初始化father
	for i := 1; i <= v; i++ {
		father[i] = i
	}
	path := make([]edge, 0)
	var v1, v2, val int
	//读取边
	for i := 0; i < e; i++ {
		fmt.Scanf("%d %d %d\n", &v1, &v2, &val)
		edges[i].V1 = v1
		edges[i].V2 = v2
		edges[i].Val = val
	}
	// 排序边
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Val < edges[j].Val
	})

	var rs int
	//遍历边
	for i := 0; i < e; i++ {
		v1 = edges[i].V1
		v2 = edges[i].V2
		if !isSame(v1, v2) {
			path = append(path, edges[i])
			join(v1, v2)
			rs += edges[i].Val
		}
	}
	fmt.Println(rs)
	for _, pathval := range path {
		fmt.Printf("%d->%d\n", pathval.V1, pathval.V2)
	}
}

func join(u, v int) {
	fu := find(u)
	fv := find(v)
	if fu == fv {
		return
	}
	father[fu] = fv
}

func find(u int) int {
	if father[u] == u {
		return u
	}
	father[u] = find(father[u])
	return father[u]
}

func isSame(u, v int) bool {
	return find(u) == find(v)
}
