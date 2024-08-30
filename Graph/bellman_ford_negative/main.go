package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type edge struct {
	from int
	to   int
	cost int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	fmt.Fscanln(r, &n, &m)
	edges := make([]edge, 0, m)
	var from, to, cost int
	for i := 0; i < m; i++ {
		fmt.Fscanln(r, &from, &to, &cost)
		edges = append(edges, edge{from, to, cost})
	}
	fmt.Println(len(edges))
	// Bellman-Ford
	// 初始化
	// 1号节点到其他节点的最短距离
	mindist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mindist[i] = math.MaxInt
	}
	mindist[1] = 0
	var flag bool
	// 张弛次数
	for i := 1; i <= n; i++ {
		for _, e := range edges {
			//更新mindist
			if i < n {
				if mindist[e.from] != math.MaxInt && mindist[e.to] > mindist[e.from]+e.cost {
					mindist[e.to] = mindist[e.from] + e.cost
				}
			} else {
				if mindist[e.from] != math.MaxInt && mindist[e.to] > mindist[e.from]+e.cost {
					flag = true
				}
			}
		}
	}
	if flag {
		fmt.Fprintln(w, "circle")
	} else if mindist[n] == math.MaxInt {
		fmt.Fprintln(w, "unconnected")
	} else {
		fmt.Fprintln(w, mindist[n])
	}

}
