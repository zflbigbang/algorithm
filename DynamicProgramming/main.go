package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func largestRectangleArea(heights []int) int {
	rs := 0
	stack := []int{0}
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	for i := 1; i < n; i++ {
		if heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					left := stack[len(stack)-1]
					h := heights[mid]
					w := i - left - 1
					rs = max(rs, h*w)
				}
			}
			stack = append(stack, i)
		}
	}
	return rs
}
func trap(height []int) int {
	rs := 0
	n := len(height)
	stack := []int{0}
	for i := 1; i < n; i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					left := stack[len(stack)-1]
					h := min(height[i], height[left]) - height[mid]
					w := i - left - 1
					rs += h * w
				}
			}
			stack = append(stack, i)
		}
	}
	return rs
}
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	rs := make([]int, n)
	for i := range rs {
		rs[i] = -1
	}
	stack := []int{0}
	for i := 1; i < len(nums)*2; i++ {
		if nums[i%len(nums)] <= nums[stack[len(stack)-1]] {
			stack = append(stack, i%len(nums))
		} else {
			for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
				rs[stack[len(stack)-1]] = nums[i%len(nums)]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i%len(nums))
		}
	}
	return rs
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := len(nums1)
	rs := make([]int, m)
	map12 := make(map[int]int, m)
	for i := 0; i < m; i++ {
		rs[i] = -1
		map12[nums1[i]] = i
	}
	n := len(nums2)
	stack := []int{0}
	for i := 1; i < n; i++ {
		if nums2[i] <= nums2[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
				if v, ok := map12[nums2[stack[len(stack)-1]]]; ok {
					rs[v] = nums2[i]
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return rs
}
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]int, n)
	rs := make([]int, n)
	stack[0] = 0
	for i := 1; i < n; i++ {
		if temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				rs[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return rs
}
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	rs := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					rs++
					dp[i][j] = true
				}
			}
		}
	}
	return rs
}
func minDistance2(word1 string, word2 string) int {
	n := len(word1) + 1
	m := len(word2) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		dp[i][0] = i
	}
	for i := 1; i < m; i++ {
		dp[0][i] = i
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[n-1][m-1]
}
func minDistance(word1 string, word2 string) int {
	n := len(word1) + 1
	m := len(word2) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		dp[i][0] = i
	}
	for i := 1; i < m; i++ {
		dp[0][i] = i
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[n-1][m-1]
}
func numDistinct(s string, t string) int {
	n := len(s) + 1
	m := len(t) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[i][0] = 1
		for j := 1; j < m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][m-1]
}
func isSubsequence(s string, t string) bool {
	n := len(s) + 1
	m := len(t) + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1] == n-1
}
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	rs := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > rs {
			rs = dp[i]
		}
	}
	return rs
}
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
func findLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	rs := 0
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > rs {
				rs = dp[i][j]
			}
		}
	}
	return rs
}
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	rs := 1
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i-1] < nums[i] {
			dp[i] = max(dp[i], dp[i-1]+1)
		}
		if dp[i] > rs {
			rs = dp[i]
		}
	}
	return rs
}
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	rs := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > rs {
			rs = dp[i]
		}
	}
	return rs
}
func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
func maxProfit5(k int, prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i < 2*k; i += 2 {
		dp[0][i] = -prices[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[n-1][2*k]
}
func maxProfit4(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][2]-prices[i]))
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = dp[i-1][1]
		dp[i][3] = max(dp[i-1][3], dp[i][2])
	}
	return max(dp[len(prices)-1][1], max(dp[len(prices)-1][2], dp[len(prices)-1][3]))
}
func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[1]
	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}
func maxProfit2(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
func maxProfit1(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}

func rob3(root *TreeNode) int {
	rs := robTree(root)
	return max(rs[0], rs[1])
}
func robTree(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val, 0}
	}
	left := robTree(root.Left)
	right := robTree(root.Right)
	val2 := max(left[0], left[1]) + max(right[0], right[1])
	val1 := root.Val + left[1] + right[1]
	return []int{val1, val2}
}
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := rob1(nums[:len(nums)-1])
	b := rob1(nums[1:])
	return max(a, b)
}
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	w := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		w[word] = true
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && w[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
func numSquares2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt
	}
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		for j := 1; j <= 2; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		num0, num1 := findStr(strs[i])
		for j := m; j >= num0; j-- {
			for k := n; k >= num1; k-- {
				dp[j][k] = max(dp[j][k], dp[j-num0][k-num1]+1)
			}
		}
	}
	return dp[m][n]
}
func findStr(str string) (num0 int, num1 int) {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			num0++
		}
	}
	num1 = len(str) - num0
	return
}
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	if target > 0 && target > sum || -target > sum {
		return 0
	}
	x := (sum + target) / 2
	dp := make([]int, x+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := x; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[x]
}
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[target]*2
}
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	bw := sum / 2
	dp := make([]int, bw+1)
	for i := 0; i < len(nums); i++ {
		for j := bw; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if dp[bw] == bw {
		return true
	} else {
		return false
	}
}
func bagProblemWithScrollingArray(w []int, v []int, bw int) int {
	dp := make([]int, bw+1)
	for i := 0; i < len(w); i++ {
		for j := bw; j >= w[i]; j-- {
			dp[j] = max(dp[j], dp[j-w[i]]+v[i])
		}
	}
	return dp[bw]
}
func bagProblem(w []int, v []int, bw int) int {
	dp := make([][]int, len(w))
	for i := 0; i < len(w); i++ {
		dp[i] = make([]int, bw+1)
	}
	for i := w[0]; i < bw+1; i++ {
		dp[0][i] = v[0]
	}
	for i := 1; i < len(w); i++ {
		for j := 1; j < bw+1; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-w[i]]+v[i])
			}
		}
	}
	return dp[len(w)-1][bw]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始第一行
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] != 0 {
			break
		}
		dp[0][i] = 1
	}
	// 初始第一列
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] != 0 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始第一行
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 初始第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := [3]int{}
	dp[1] = 1
	dp[2] = 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = dp[1] + dp[2]
		dp[1] = dp[2]
		dp[2] = sum
	}
	return sum
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := [2]int{}
	dp[0] = 0
	dp[1] = 1
	sum := 0
	for i := 1; i < n; i++ {
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return sum
}
func integerBreak(n int) int {
	// dp[i] 表示 i 拆分成数乘积最大
	dp := make([]int, n+1)
	// 初始化
	dp[2] = 1
	// 遍历 + 状态转移
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(max(j*(i-j), j*dp[i-j]), dp[i])
		}
	}
	return dp[n]
}
