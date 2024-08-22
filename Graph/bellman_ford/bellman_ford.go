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
	var g = make([]edge, 0, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		g = append(g, edge{a, b, c})
	}
	var minDist = make([]int, n+1)
	for i := 1; i <= n; i++ {
		minDist[i] = math.MaxInt
	}
	minDist[1] = 0
	//遍历 n - 1 次
	for i := 1; i < n; i++ {
		//遍历所有边
		for _, e := range g {
			// 松弛
			if minDist[e.from] != math.MaxInt && minDist[e.to] > minDist[e.from]+e.cost {
				minDist[e.to] = minDist[e.from] + e.cost
			}
		}

	}
	if minDist[n] == math.MaxInt {
		fmt.Fprintln(w, "unconnected")
	} else {
		fmt.Fprintln(w, minDist[n])
	}
}
