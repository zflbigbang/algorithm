package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// 最小堆
type Data struct {
	x, y    int
	g, h, f int
}

var endX, endY int

// 移动的步数
var moves [1001][1001]int

// 移动的方向
var dir = [8][2]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}

type MinHeap []Data

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].f < h[j].f

}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Data))
}

// 欧式距离
func Heu(d Data) int {
	return (d.x-endX)*(d.x-endX) + (d.y-endY)*(d.y-endY)
}

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(r, &n)
	defer w.Flush()
	var startX, startY int
	var d Data
	for i := 0; i < n; i++ {
		//重置 moves
		moves = [1001][1001]int{}
		fmt.Fscanln(r, &startX, &startY, &endX, &endY)
		d = Data{startX, startY, 0, 0, 0}
		d.h = Heu(d)
		d.f = d.h + d.g
		astar(d)
		fmt.Fprintln(w, moves[endX][endY])
	}

}

func astar(d Data) {
	var cur, next Data
	var minHeap MinHeap
	heap.Push(&minHeap, d)
	for minHeap.Len() > 0 {
		cur = heap.Pop(&minHeap).(Data)
		if cur.x == endX && cur.y == endY {
			return
		}
		for _, v := range dir {
			nextX := cur.x + v[0]
			nextY := cur.y + v[1]
			if nextX < 1 || nextX > 1000 || nextY < 1 || nextY > 1000 || moves[nextX][nextY] != 0 {
				continue
			}
			h := Heu(next)
			f := h + cur.g + 5
			next = Data{nextX, nextY, cur.g + 5, h, f}
			moves[next.x][next.y] = moves[cur.x][cur.y] + 1
			heap.Push(&minHeap, next)
		}
	}
}
