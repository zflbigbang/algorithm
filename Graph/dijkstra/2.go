package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge 表示带权重的边
type Edge struct {
	to, val int
}

// PriorityQueue 实现一个小顶堆
type Item struct {
	node, dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(n, m int, edges [][]int, start, end int) int {
	grid := make([][]Edge, n+1)
	for _, edge := range edges {
		p1, p2, val := edge[0], edge[1], edge[2]
		grid[p1] = append(grid[p1], Edge{to: p2, val: val})
	}

	minDist := make([]int, n+1)
	for i := range minDist {
		minDist[i] = math.MaxInt64
	}
	visited := make([]bool, n+1)

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, dist: 0})
	minDist[start] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)

		if visited[cur.node] {
			continue
		}

		visited[cur.node] = true

		for _, edge := range grid[cur.node] {
			if !visited[edge.to] && minDist[cur.node]+edge.val < minDist[edge.to] {
				minDist[edge.to] = minDist[cur.node] + edge.val
				heap.Push(pq, &Item{node: edge.to, dist: minDist[edge.to]})
			}
		}
	}

	if minDist[end] == math.MaxInt64 {
		return -1
	}
	return minDist[end]
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var p1, p2, val int
		fmt.Scan(&p1, &p2, &val)
		edges[i] = []int{p1, p2, val}
	}

	start := 1 // 起点
	end := n   // 终点

	result := dijkstra(n, m, edges, start, end)
	fmt.Println(result)
}
