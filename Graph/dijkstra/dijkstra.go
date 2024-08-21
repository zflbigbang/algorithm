package main

import (
	"fmt"
)

func main1() {
	//选取距离源点最短的节点，
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	grid := make([][]int, n+1)
	visited := make([]bool, n+1)
	//记录每个节点距离源的最短距离
	minDist := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		grid[i] = make([]int, n+1)
		minDist[i] = 10001
	}
	//初始化边
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			grid[i][j] = 10001
		}
	}
	var v1, v2, val int
	for i := 1; i < m+1; i++ {
		fmt.Scanf("%d %d %d\n", &v1, &v2, &val)
		grid[v1][v2] = val
	}
	minDist[1] = 0
	var cur, minVal int
	// 遍历图
	for i := 1; i < n+1; i++ {
		cur = 0
		minVal = 10001
		//选择最短距离的点且没有访问过的
		for j := 1; j < n+1; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				cur = j
			}
		}
		// 标记访问,同时更新与该节点相连的距离
		visited[cur] = true
		for k, x := range grid[cur] {
			// 如果小于，则更新
			if !visited[k] && x != 10001 && minDist[cur]+x < minDist[k] {
				minDist[k] = minDist[cur] + x
			}
		}
	}
	if minDist[n] != 10001 {
		fmt.Print(minDist[n])
		return
	}
	fmt.Print(-1)

}
