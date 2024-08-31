package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type edge struct {
	from, to, cost int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m int
	_, _ = fmt.Fscanln(r, &n, &m)
	edges := make([]edge, m)
	var from, to, cost int
	for i := 0; i < m; i++ {
		fmt.Fscanln(r, &from, &to, &cost)
		edges[i] = edge{from, to, cost}
	}
	var src, dist, k int
	fmt.Fscanln(r, &src, &dist, &k)
	// bellman-ford
	minDist := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		minDist[i] = math.MaxInt
	}
	minDist[src] = 0
	pre := make([]int, n+1)
	for i := 1; i <= k+1; i++ {
		copy(pre, minDist)
		for _, e := range edges {
			if pre[e.from] != math.MaxInt && pre[e.to] > pre[e.from]+e.cost {
				minDist[e.to] = pre[e.from] + e.cost
			}
		}
	}
	if minDist[dist] == math.MaxInt {
		fmt.Fprintln(w, "unreachable")
	} else {
		fmt.Fprintln(w, minDist[dist])
	}
}
