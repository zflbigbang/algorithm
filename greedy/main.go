package main

import (
	"container/list"
	"math"
	"sort"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	rs := make([][]int, 0, len(intervals))
	rs = append(rs, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= rs[len(rs)-1][1] {
			rs[len(rs)-1][1] = max(intervals[i][1], rs[len(rs)-1][1])
		} else {
			rs = append(rs, intervals[i])
		}
	}
	return rs
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func partitionLabels(s string) []int {
	b := []byte(s)
	hash := [26]int{}

	rs := make([]int, 0)
	for i := 0; i < len(b); i++ {
		hash[b[i]-'0'] = i
	}
	left := 0
	right := 0
	for i := 0; i < len(b); i++ {
		right = max(hash[b[i]-'0'], right)
		if right == i {
			rs = append(rs, right-left+1)
			left = i + 1
		}
	}
	return rs
}
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	rs := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
			rs++
		}
	}
	return rs
}
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	rs := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			rs++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}
	return rs
}

func monotoneIncreasingDigits(n int) int {
	t := strconv.Itoa(n)
	s := []byte(t)
	flag := len(s)
	for i := len(s) - 1; i > 0; i-- {
		if s[i] < s[i-1] {
			s[i-1]--
			flag = i
		}
	}
	for i := flag; i < len(s); i++ {
		s[i] = '9'
	}
	rs, _ := strconv.Atoi(string(s))
	return rs
}
func minCameraCover(root *TreeNode) int {
	rs := 0
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := travel(root.Left)
		right := travel(root.Right)
		if left == 2 && right == 2 {
			return 0
		} else if left == 0 || right == 0 {
			rs++
			return 1
		} else {
			return 2
		}
	}
	if travel(root) == 0 {
		rs++
	}
	return rs
}
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		} else {
			if five > 0 && ten > 0 {
				ten--
				five--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totSum := 0
	rs := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totSum += gas[i] - cost[i]
		if curSum < 0 {
			curSum = 0
			rs = i + 1
		}
	}
	if totSum >= 0 {
		return rs
	}
	return -1
}
func candy(ratings []int) int {
	rs := 0
	cs := make([]int, len(ratings))
	for i := 0; i < len(cs); i++ {
		cs[i] = 1
	}
	//right > left
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			cs[i+1] = cs[i] + 1
		}
	}
	// left > right
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			cs[i-1] = max(cs[i]+1, cs[i-1])
		}
	}
	for i := 0; i < len(cs); i++ {
		rs += cs[i]
	}
	return rs
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	//复制效率比较慢，所有切片向后移动
	//for i, person := range people {
	//	copy(people[person[1]+1:i+1], people[person[1]:i])
	//	people[person[1]] = person
	//}
	//return people
	//使用链表
	l := list.New() //创建链表
	for i := 0; i < len(people); i++ {
		position := people[i][1]
		mark := l.PushBack(people[i]) //插入元素
		e := l.Front()
		for position != 0 { //获取相对位置
			position--
			e = e.Next()
		}
		l.MoveBefore(mark, e) //移动位置

	}
	res := [][]int{}
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}
func largestSumAfterKNegations(nums []int, k int) int {
	rs := 0
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	for i := 0; i < len(nums); i++ {
		rs += nums[i]
	}
	return rs
}
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cur := 0
	rs := 0
	next := 0
	for i := 0; i <= len(nums); i++ {
		if i+nums[i] > next {
			next = i + nums[i]
		}
		if i == cur {
			cur = next
			rs++
			if cur >= len(nums)-1 {
				break
			}
		}
	}
	return rs
}
func canJump(nums []int) bool {
	cov := nums[0]
	for i := 0; i <= cov; i++ {
		if i+nums[i] > cov {
			cov = i + nums[i]
		}
		if cov >= len(nums)-1 {
			return true
		}
	}
	return false
}
func maxProfit(prices []int) int {
	rs := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			rs += prices[i+1] - prices[i]
		}
	}
	return rs
}
func maxSubArray(nums []int) int {
	rs := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > rs {
			rs = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return rs
}
func wiggleMaxLength(nums []int) int {
	preDiff := 0
	curDiff := 0
	rs := 0
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if preDiff >= 0 && curDiff < 0 || preDiff <= 0 && curDiff > 0 {
			rs++
			preDiff = curDiff
		}
	}
	return rs
}
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && s[i] >= g[index] {
			index++
		}
	}
	return index
}
