package main

import (
	"sort"
)

func main() {

}

func solveSudoku(board [][]byte) {
	var backtracking func() bool
	backtracking = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					for k := '1'; k <= '9'; k++ {
						if isValidS(byte(k), i, j, board) {
							board[i][j] = byte(k)
							if backtracking() {
								return true
							}
							board[i][j] = '.'
						}
					}
					return false
				}
			}
		}
		return true
	}
	backtracking()
}
func isValidS(k byte, row, col int, board [][]byte) bool {
	// row
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	// col
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}
	// 3*3
	row = (row / 3) * 3
	col = (col / 3) * 3
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
func solveNQueens(n int) [][]string {
	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}
	var rs [][]string
	var backtracking func(row int)
	backtracking = func(row int) {
		if row == n {
			tmp := make([]string, 0, n)
			for _, bytes := range path {
				tmp = append(tmp, string(bytes))
			}
			rs = append(rs, tmp)
		}
		for i := 0; i < n; i++ {
			if isValid(n, i, row, path) {
				path[row][i] = 'Q'
				backtracking(row + 1)
				path[row][i] = '.'
			}
		}
	}
	backtracking(0)
	return rs
}
func isValid(n int, col int, row int, path [][]byte) bool {
	//check col
	for i := 0; i < row; i++ {
		if path[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	//check 45°
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	return true
}
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var rs [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				backtracking()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	backtracking()
	return rs
}
func permute(nums []int) [][]int {
	var rs [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				path = append(path, nums[i])
				used[i] = true
				backtracking()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	backtracking()
	return rs
}
func findSubsequences(nums []int) [][]int {
	var rs [][]int
	var path []int
	var backtracking func(st int)
	backtracking = func(st int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}
		used := make(map[int]bool, len(nums)-st)
		for i := st; i < len(nums); i++ {
			if used[nums[i]] || len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return rs
}
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var rs [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtracking func(st int)
	backtracking = func(st int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		rs = append(rs, tmp)
		for i := st; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtracking(i + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking(0)
	return rs
}
func subsets(nums []int) [][]int {
	var rs [][]int
	var path []int
	var backtracking func(st int)
	backtracking = func(st int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		rs = append(rs, tmp)
		for i := st; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return rs
}

//	func restoreIpAddresses(s string) []string {
//		var rs []string
//		var path []string
//		var backtracking func(st int)
//		backtracking = func(st int) {
//			if len(path) == 4 {
//				if st == len(s) {
//					str := strings.Join(path, ".")
//					rs = append(rs, str)
//				}
//				return
//			}
//			for i := st; i < len(s); i++ {
//				str := s[st : i+1]
//				if isValid(str) {
//					path = append(path, str+".")
//					backtracking(i + 1)
//					path = path[:len(path)-1]
//				} else {
//					break
//				}
//			}
//		}
//		backtracking(0)
//		return rs
//	}
//
//	func isValid(str string) bool {
//		if str[0] == '0' && len(str) > 1 {
//			return false
//		}
//		num, _ := strconv.Atoi(str)
//		if num < 0 || num > 255 {
//			return false
//		}
//		return true
//	}
func partition(s string) [][]string {
	var rs [][]string
	var path []string
	var backtracking func(st int)
	backtracking = func(st int) {
		if st == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}
		for i := st; i < len(s); i++ {
			if !isPalindrome(s, st, i) {
				continue
			}
			str := s[st : i+1]
			path = append(path, str)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return rs
}
func isPalindrome(s string, st int, end int) bool {
	for st < end {
		if s[st] != s[end] {
			return false
		}
		st++
		end--
	}
	return true
}
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var rs [][]int
	path := make([]int, len(candidates))
	used := make([]bool, len(candidates))
	var backtracking func(st int, sum int)
	backtracking = func(st int, sum int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}
		for i := st; i < len(candidates) && sum+candidates[i] <= target; i++ {
			if candidates[i] == candidates[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, candidates[i])
			backtracking(i+1, sum+candidates[i])
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking(0, 0)
	return rs
}
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var rs [][]int
	var path []int
	var backtracking func(st int, sum int)
	backtracking = func(st int, sum int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			rs = append(rs, tmp)
		}
		for i := st; i < len(candidates) && sum+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			backtracking(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return rs
}
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var rs []string
	var path []byte
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtracking func(st int)
	backtracking = func(st int) {
		if st == len(digits) {
			tmp := string(path)
			rs = append(rs, tmp)
			return
		}
		index := digits[st] - '0'
		s := m[index]
		for i := 0; i < len(s); i++ {
			path = append(path, s[i])
			backtracking(st + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return rs
}
func combinationSum3(k int, n int) [][]int {
	var rs [][]int
	var path []int
	var backtracking func(st int, sum int)
	backtracking = func(st int, sum int) {
		if sum > n {
			return
		}
		if len(path) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				rs = append(rs, tmp)
				return
			} else {
				return
			}
		}
		for i := st; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			sum += i
			backtracking(i+1, sum)
			sum -= i
			path = path[:len(path)-1]
		}
	}
	backtracking(1, 0)
	return rs
}
func combine(n int, k int) [][]int {
	var rs [][]int
	var path []int
	var backtracking func(st int)
	backtracking = func(st int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			rs = append(rs, tmp)
			return
		}
		for i := st; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return rs
}
