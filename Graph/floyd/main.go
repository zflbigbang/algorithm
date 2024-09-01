package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, m int
	defer w.Flush()
	_, _ = fmt.Fscanln(r, &n, &m)
	const INF = 1e5
	//定义和初始化dp
	dp := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			dp[i][j] = INF

		}
	}

	var u, v, cost int
	for i := 0; i < m; i++ {
		fmt.Fscanln(r, &u, &v, &cost)
		dp[u][v] = cost
		dp[v][u] = cost
	}

	//开始floyd
	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				if dp[i][k] < INF && dp[k][j] < INF {
					dp[i][j] = Min(dp[i][j], dp[i][k]+dp[k][j])
				}

			}
		}
	}
	var q int
	fmt.Fscanln(r, &q)
	for i := 0; i < q; i++ {
		fmt.Fscanln(r, &u, &v)
		if dp[u][v] == INF {
			fmt.Fprintln(w, -1)
		} else {
			fmt.Fprintln(w, dp[u][v])
		}
	}
}

// 辅助函数：求最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
