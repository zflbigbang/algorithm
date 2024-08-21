package main

import "fmt"

func main() {
	//选取入度为零的节点放入队列，如果没有，则无法拓扑排序 bfs 是核心
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	//定义每个节点的入度
	inDegree := make([]int, n)
	rs := make([]int, 0, n)
	//模拟队列
	queue := make([]int, 0, n)
	//记录每个节点指向的所有的节点
	toNode := make(map[int][]int)
	var v1, v2 int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &v1, &v2)
		inDegree[v2]++
		toNode[v1] = append(toNode[v1], v2)
	}
	for i := 0; i < n; i++ {
		//选择入度为0节点加入队列
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	//遍历queue
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		//选择入度为0节点加入队列
		if inDegree[cur] == 0 {
			rs = append(rs, cur)
			//当节点加入队列时，更新这个节点指向所有节点的入度
			for _, v := range toNode[cur] {
				inDegree[v]--
				if inDegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
	}
	//打印结果
	if n == len(rs) {
		for i := 0; i < n-1; i++ {
			fmt.Printf("%d ", rs[i])

		}
		fmt.Print(rs[n-1])
	} else {
		fmt.Print(-1)
	}
}
