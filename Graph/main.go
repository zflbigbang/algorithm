package main

func main() {
	largestIsland([][]int{{1, 1}, {1, 1}})
}

func findCircleNum(isConnected [][]int) int {
	cnt := 0
	n := len(isConnected)
	visited := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		visited[cur] = true
		for to, conn := range isConnected[cur] {
			if conn == 1 && !visited[to] {
				dfs(to)
			}
		}
	}
	for i, v := range visited {
		if !v {
			cnt++
			dfs(i)
		}
	}

	return cnt

}

func largestIsland(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	rs := 0
	cnt := 0
	mark := 2
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	markArea := make(map[int]int, 0)
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		if visited[x][y] || grid[x][y] == 0 {
			return
		}
		cnt++
		visited[x][y] = true
		grid[x][y] = mark
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx < 0 || nextx >= n || nexty < 0 || nexty >= m {
				continue
			}
			dfs(nextx, nexty)
		}
	}
	isAllgrid := true
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				isAllgrid = false
			}
			if !visited[i][j] && grid[i][j] == 1 {
				cnt = 0
				dfs(i, j)
				markArea[mark] = cnt
				mark++
			}
		}
	}
	if isAllgrid {
		return cnt
	}
	//遍历图中的零，改零为一，计算面积
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				cnt = 1
				visitedArea := make(map[int]bool, 0)
				for k := 0; k < 4; k++ {
					nextx := i + dir[k][0]
					nexty := j + dir[k][1]
					if nextx < 0 || nextx >= n || nexty < 0 || nexty >= m {
						continue
					}
					if visitedArea[grid[nextx][nexty]] {
						continue
					}
					cnt += markArea[grid[nextx][nexty]]
					visitedArea[grid[nextx][nexty]] = true
				}
			}
			if cnt > rs {
				rs = cnt
			}
		}
	}
	return rs
}
func pacificAtlantic(heights [][]int) [][]int {
	n := len(heights)
	m := len(heights[0])
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pacific := make([][]bool, n)
	atlantic := make([][]bool, n)
	rs := make([][]int, 0)
	for i := range pacific {
		pacific[i] = make([]bool, m)
		atlantic[i] = make([]bool, m)

	}
	var dfs func(x int, y int, visited [][]bool)
	dfs = func(x int, y int, visited [][]bool) {
		if visited[x][y] {
			return
		}
		t := heights[x][y]
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx < 0 || nextx >= n || nexty < 0 || nexty >= m {
				continue
			}
			if heights[nextx][nexty] < t {
				continue
			}
			dfs(nextx, nexty, visited)
		}
	}
	//左右两侧
	for i := 0; i < n; i++ {
		if !pacific[i][0] {
			dfs(i, 0, pacific)
		}
		if !atlantic[i][m-1] {
			dfs(i, m-1, atlantic)
		}
	}
	//上下两侧
	for i := 0; i < m; i++ {
		if !pacific[0][i] {
			dfs(0, i, pacific)
		}
		if !atlantic[n-1][i] {
			dfs(n-1, i, atlantic)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i][j] && atlantic[i][j] {
				rs = append(rs, []int{i, j})
			}
		}
	}
	return rs
}
func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx < 0 || nextx >= n || nexty < 0 || nexty >= m {
				continue
			}
			if board[nextx][nexty] == 'O' {
				dfs(nextx, nexty)
			}
		}
	}
	//左右两侧
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][m-1] == 'O' {
			dfs(i, m-1)
		}
	}
	//上下两侧
	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[n-1][i] == 'O' {
			dfs(n-1, i)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
func numEnclaves(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	cnt := 0
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		cnt++
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
				continue
			}
			if !visited[nextx][nexty] && grid[nextx][nexty] == 1 {
				visited[nextx][nexty] = true
				dfs(nextx, nexty)
			}
		}
	}
	// 左右两侧
	for i := 0; i < len(grid); i++ {
		if !visited[i][0] && grid[i][0] == 1 {
			visited[i][0] = true
			dfs(i, 0)
		}
		if !visited[i][len(grid[0])-1] && grid[i][len(grid[0])-1] == 1 {
			visited[i][len(grid[0])-1] = true
			dfs(i, len(grid[0])-1)
		}
	}
	// 上下两侧
	for i := 0; i < len(grid[0]); i++ {
		if !visited[0][i] && grid[0][i] == 1 {
			visited[0][i] = true
			dfs(0, i)
		}
		if !visited[len(grid)-1][i] && grid[len(grid)-1][i] == 1 {
			visited[len(grid)-1][i] = true
			dfs(len(grid)-1, i)
		}
	}
	cnt = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}
func numEnclaves1(grid [][]int) int {
	cnt := 0
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		cnt++
		grid[x][y] = 0
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
				continue
			}
			if grid[nextx][nexty] == 1 {
				dfs(nextx, nexty)
			}
		}
	}
	// 左右两侧
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
			dfs(i, 0)
		}
		if grid[i][len(grid[0])-1] == 1 {
			dfs(i, len(grid[0])-1)
		}
	}
	// 上下两侧
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 1 {
			dfs(0, i)
		}
		if grid[len(grid)-1][i] == 1 {
			dfs(len(grid)-1, i)
		}
	}
	cnt = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}
func maxAreaOfIsland1(grid [][]int) int {
	rs := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	cnt := 0
	queue := make([][]int, 0)
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var bfs func(x int, y int)
	bfs = func(x int, y int) {
		// 入队和标记
		queue = append(queue, []int{x, y})
		visited[x][y] = true
		for len(queue) > 0 {
			// 出队
			curx := queue[len(queue)-1][0]
			cury := queue[len(queue)-1][1]
			queue = queue[:len(queue)-1]
			//四个方向遍历
			for i := 0; i < 4; i++ {
				nextx := dir[i][0] + curx
				nexty := dir[i][1] + cury
				if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
					continue
				}
				if !visited[nextx][nexty] && grid[nextx][nexty] == 1 {
					//入队和标记,计数
					queue = append(queue, []int{nextx, nexty})
					cnt++
					visited[nextx][nexty] = true
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				cnt = 1
				bfs(i, j)
				if cnt > rs {
					rs = cnt
				}
			}
		}
	}
	return rs
}
func maxAreaOfIsland(grid [][]int) int {
	rs := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cnt := 0
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		for i := 0; i < 4; i++ {
			nextx := x + dir[i][0]
			nexty := y + dir[i][1]
			if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
				continue
			}
			if !visited[nextx][nexty] && grid[nextx][nexty] == 1 {
				visited[nextx][nexty] = true
				cnt++
				dfs(nextx, nexty)
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				cnt = 1
				visited[i][j] = true
				dfs(i, j)
				if cnt > rs {
					rs = cnt
				}
			}
		}
	}
	return rs
}
func numIslands1(grid [][]byte) int {
	rs := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		for i := 0; i < 4; i++ {
			nextx := dir[i][0] + x
			nexty := dir[i][1] + y
			if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
				continue
			}
			if !visited[nextx][nexty] && grid[nextx][nexty] == '1' {
				//标记和递归
				visited[nextx][nexty] = true
				dfs(nextx, nexty)
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				rs++
				dfs(i, j)
			}
		}
	}
	return rs
}
func numIslands(grid [][]byte) int {
	rs := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	queue := make([][]int, 0)
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var bfs func(x int, y int)
	bfs = func(x int, y int) {
		// 入队和标记
		queue = append(queue, []int{x, y})
		visited[x][y] = true
		for len(queue) > 0 {
			// 出队
			curx := queue[len(queue)-1][0]
			cury := queue[len(queue)-1][1]
			queue = queue[:len(queue)-1]
			//四个方向遍历
			for i := 0; i < 4; i++ {
				nextx := dir[i][0] + curx
				nexty := dir[i][1] + cury
				if nextx < 0 || nextx >= len(grid) || nexty < 0 || nexty >= len(grid[0]) {
					continue
				}
				if !visited[nextx][nexty] && grid[nextx][nexty] == '1' {
					//入队和标记
					queue = append(queue, []int{nextx, nexty})
					visited[nextx][nexty] = true
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				rs++
				bfs(i, j)
			}
		}
	}
	return rs
}
func allPathsSourceTarget(graph [][]int) [][]int {
	rs := make([][]int, 0)
	path := []int{0}
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
			return
		}
		for i := 0; i < len(graph[x]); i++ {
			path = append(path, graph[x][i])
			dfs(graph[x][i])
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return rs
}
