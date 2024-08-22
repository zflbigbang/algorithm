package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 6 5
// 5 6 1
// 4 5 1
// 3 4 1
// 2 3 1
// 1 2 1
type edge struct {
	from, to, cost int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	fmt.Fscan(r, &n, &m)
	var a, b, c int
	var g = make([][]edge, 0)
	for i := 0; i <= n; i++ {
		g = append(g, make([]edge, 0))
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		g[a] = append(g[a], edge{from: a, to: b, cost: c})
	}
	var minDist = make([]int, n+1)
	for i := 1; i <= n; i++ {
		minDist[i] = math.MaxInt
	}
	minDist[1] = 0
	que := make([]int, 0, n)
	isVisited := make([]bool, n+1)
	//入队
	que = append(que, 1)
	for len(que) > 0 {
		//出队
		u := que[0]
		que = que[1:]
		isVisited[u] = false
		for _, e := range g[u] {
			if minDist[e.to] > minDist[u]+e.cost {
				minDist[e.to] = minDist[u] + e.cost
				if !isVisited[e.to] {
					que = append(que, e.to)
					isVisited[e.to] = true
				}
			}
		}
	}
	if minDist[n] == math.MaxInt {
		fmt.Fprintln(w, "unconnected")
	} else {
		fmt.Fprintln(w, minDist[n])
	}
}
