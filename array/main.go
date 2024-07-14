package main

import (
	"math"
)

func main() {

}

func countPrimes(n int) (cnt int) {
	// 初始化切片，长度为n，所有元素默认为true
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	rs := make([]int, n*m)
	cnt := max(n, m)
	loop := cnt / 2
	i, j, x, y := 0, 0, 0, 0
	offset, t := 1, 0
	for loop > 0 {
		for j = y; j < m-offset; j++ {
			rs[t] = matrix[x][j]
			t++
		}
		for i = x; i < n-offset; i++ {
			rs[t] = matrix[i][j]
			t++
		}
		for ; j > y; j-- {
			rs[t] = matrix[i][j]
			t++
			if t >= n*m {
				goto here
			}
		}
		for ; i > x; i-- {
			rs[t] = matrix[i][j]
			t++
			if t >= n*m {
				goto here
			}
		}
		x++
		y++
		offset++
		loop--
	}
here:
	if cnt%2 == 1 && t < n*m {
		rs[t] = matrix[x][y]
	}
	return rs
}
func max(n int, m int) int {
	if n > m {
		return n
	} else {
		return m
	}
}
func generateMatrix(n int) [][]int {
	rs := make([][]int, n)
	for i := 0; i < n; i++ {
		rs[i] = make([]int, n)
	}
	loop := n / 2
	x, y := 0, 0
	offset := 1
	t := 0
	i, j := 0, 0
	for loop > 0 {
		for j = y; j < n-offset; j++ {
			t++
			rs[x][j] = t
		}
		for i = x; i < n-offset; i++ {
			t++
			rs[i][j] = t
		}
		for ; j > y; j-- {
			t++
			rs[i][j] = t
		}
		for ; i > x; i-- {
			t++
			rs[i][j] = t
		}
		x++
		y++
		offset++
		loop--
	}
	if n%2 == 1 {
		t++
		rs[x][y] = t
	}
	return rs
}
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	rs := math.MaxInt
	anl, anr := -1, -1
	l, r := 0, 0
	for ; r < len(s); r++ {
		if ori[s[r]] > 0 {
			break
		}
	}
	for l = r; r < len(s); r++ {
		cnt[s[r]]++
		for check() && l <= r {
			tmp := r - l + 1
			if tmp < rs {
				anl, anr = l, r
				rs = tmp
			}
			if ori[s[l]] > 0 {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if anl == -1 {
		return ""
	}
	return s[anl : anr+1]
}
func totalFruit(fruits []int) int {
	i := 0
	cnt := map[int]int{}
	rs := math.MaxInt
	for j := 0; j < len(fruits); j++ {
		cnt[fruits[j]]++
		for len(cnt) > 2 {
			cnt[fruits[i]]--
			if cnt[fruits[i]] == 0 {
				delete(cnt, fruits[i])
			}
			i++
		}
		tem := j - i + 1
		if tem < rs {
			rs = tem
		}
	}
	if rs == math.MaxInt {
		return 0
	}
	return rs
}
func minSubArrayLen(target int, nums []int) int {
	i, sum := 0, 0
	rs := math.MaxInt
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			sum -= nums[i]
			tem := j - i + 1
			if tem < rs {
				rs = tem
			}
			i++
		}
	}
	if rs == math.MaxInt {
		return 0
	}
	return rs
}
func backspaceCompare(s string, t string) bool {
	skips, skipt := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				i--
				skips++
			} else if skips > 0 {
				i--
				skips--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				j--
				skipt++
			} else if skipt > 0 {
				j--
				skipt--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
func moveZeroes(nums []int) {
	s := 0
	for f := 0; f < len(nums); f++ {
		if nums[f] != 0 {
			nums[s] = nums[f]
			s++
		}
	}
	for ; s < len(nums); s++ {
		nums[s] = 0
	}
}
func removeDuplicates(nums []int) int {
	s := 0
	for f := 1; f < len(nums); f++ {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}
	}
	return s + 1
}

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		m := (r-l)/2 + l
		if m*m < num {
			l = m + 1
		} else if m*m > num {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}
func mySqrt(x int) int {
	l, r := 0, x
	ans := 0
	for l <= r {
		m := (r-l)/2 + l
		if m*m <= x {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}

func sortedSquares(nums []int) []int {
	rs := make([]int, len(nums))
	p := len(nums) - 1
	for l, r := 0, p; l <= r; {
		x, y := nums[l]*nums[l], nums[r]*nums[r]
		if x > y {
			rs[p] = x
			l++
		} else {
			rs[p] = y
			r--
		}
		p--
	}
	return rs
}
func removeElement(nums []int, val int) int {
	f, s := 0, 0
	for ; f < len(nums); f++ {
		if nums[f] != val {
			nums[s] = nums[f]
			s++
		}
	}
	return s
}
func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	m := 0
	for l <= r {
		m = (r-l)/2 + l
		if nums[m] == target {
			goto breakHere
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return []int{-1, -1}
breakHere:
	for l = m; l >= 0; l-- {
		if nums[l] != target {
			break
		}
	}
	for r = m; r <= len(nums)-1; r++ {
		if nums[r] != target {
			break
		}
	}
	return []int{l + 1, r - 1}
}
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (r-l)/2 + l
		if nums[m] == target {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

// 二分查找
// 1.左闭右闭
func search1(arr []int, t int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := (r-l)/2 + l
		if arr[m] == t {
			return m
		} else if t < arr[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

// 2.左闭右开
func search2(arr []int, t int) int {
	l := 0
	r := len(arr)
	for l < r {
		m := (r-l)/2 + l
		if arr[m] == t {
			return m
		} else if t < arr[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return -1
}
