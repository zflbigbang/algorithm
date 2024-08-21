package main

import (
	"bufio"
	"fmt"
	"os"
)

// 初始化
var father []int

func Init(n int) {
	father = make([]int, n+1)
	size = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		father[i] = i
		size[i] = 1
	}
}

// 记录并查集的大小
var size []int

// 并
func join(a, b int) {
	fa := find(a)
	fb := find(b)
	if fa != fb {
		if size[fa] > size[fb] {
			father[fb] = fa
			size[fa] += size[fb]
		} else {
			father[fa] = fb
			size[fb] += size[fa]
		}
	}
}

// 查
func find(a int) int {
	if father[a] == a {
		return a
	} else {
		return find(father[a])
	}
}

// 是否联通
func isConnected(a, b int) bool {
	return find(a) == find(b)
}
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, m, q int
	fmt.Fscan(r, &n, &m, &q)
	Init(n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b)
		join(a, b)
	}
	var c, d, e int
	for i := 0; i < q; i++ {
		fmt.Fscan(r, &c, &d, &e)
		if c == 1 {
			if isConnected(d, e) {
				if size[d] > size[e] {
					father[e] = e
					size[d] -= size[e]
				} else {
					father[d] = d
					size[e] -= size[d]
				}
			}
		} else {
			if isConnected(d, e) {
				fmt.Fprintln(w, "Yes")
			} else {
				fmt.Fprintln(w, "No")
			}
		}
	}

}
