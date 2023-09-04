package main

import (
	"container/list"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Right = travel(root.Right)
		root.Val += pre
		pre = root.Val
		root.Left = travel(root.Left)
		return root
	}
	return travel(root)
}

func sortedArrayToBST(nums []int) *TreeNode {
	var travel func(left int, right int) *TreeNode
	travel = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (right-left)/2 + left
		root := &TreeNode{Val: nums[mid]}
		root.Left = travel(left, mid-1)
		root.Right = travel(mid+1, right)
		return root
	}
	return travel(0, len(nums)-1)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Right != nil && root.Left == nil {
			return root.Right
		} else if root.Right == nil && root.Left == nil {
			return nil
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
func lowestCommonAncestorMin(root, p, q *TreeNode) *TreeNode {
	var leftNode *TreeNode
	var rightNode *TreeNode
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val > p.Val && root.Val > q.Val {
			leftNode = travel(root.Left)
		}
		if root.Val < p.Val && root.Val < q.Val {
			rightNode = travel(root.Right)
		}
		if leftNode != nil {
			return leftNode
		}
		if rightNode != nil {
			return rightNode
		}
		return root
	}
	return travel(root)
}
func findMode(root *TreeNode) []int {
	array := make([]int, 0)
	maxCnt := 0
	cnt := 0
	var pre *TreeNode
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		if pre != nil && pre.Val == root.Val {
			cnt++
		} else {
			cnt = 1
		}
		pre = root
		if cnt == maxCnt {
			array = append(array, root.Val)
		} else if cnt > maxCnt {
			maxCnt = cnt
			array = []int{}
			array = append(array, root.Val)
		}
		travel(root.Right)
	}
	travel(root)
	return array
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	} else if leftNode != nil && rightNode == nil {
		return leftNode
	} else if leftNode == nil && rightNode != nil {
		return rightNode
	}
	return nil
}
func getMinimumDifference(root *TreeNode) int {
	minRs := math.MaxInt
	var pre *TreeNode
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		if pre != nil && root.Val-pre.Val < minRs {
			minRs = root.Val - pre.Val
		}
		pre = root
		travel(root.Right)
	}
	travel(root)
	return minRs
}
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var travel func(root *TreeNode) bool
	travel = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		leftFlag := travel(root.Left)
		if pre != nil && pre.Val > root.Val {
			return false
		}
		pre = root
		rightFlag := travel(root.Right)
		return leftFlag && rightFlag
	}
	return travel(root)
}
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
func findMax(nums []int) (index int) {
	maxTmp := 0
	for i, num := range nums {
		if num > maxTmp {
			maxTmp = num
			index = i
		}
	}
	return
}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	//终止条件
	if len(nums) == 1 {
		root := &TreeNode{Val: nums[0]}
		return root
	}
	//寻找最大值下标
	index := findMax(nums)
	root := &TreeNode{Val: nums[index]}
	//单层递归逻辑
	//向左切割数组并递归
	if index > 0 {
		root.Left = constructMaximumBinaryTree(nums[:index])
	}
	if index < len(nums)-1 {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}
	return root
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	tmp := postorder[len(postorder)]
	root := &TreeNode{Val: tmp}
	if len(postorder) == 1 {
		return root
	}
	//寻找切割点
	index := 0
	for ; index < len(postorder); index++ {
		if inorder[index] == tmp {
			break
		}
	}
	//切割前序数组
	inorderLeft := inorder[:index]
	inorderRight := inorder[index+1:]
	//切割后序数组
	postorderLeft := postorder[:index]
	postorderRight := postorder[:len(postorder)-1]
	root.Left = buildTree(inorderLeft, postorderLeft)
	root.Right = buildTree(inorderRight, postorderRight)
	return root
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

var rs int
var maxD int

func findBottomLeftValue(root *TreeNode) int {
	rs, maxD = 0, 0
	dfsFindBottomLeftValue(root, 0)
	return rs
}
func dfsFindBottomLeftValue(root *TreeNode, d int) {
	if root == nil {
		return
	}
	if d > maxD {
		maxD = d
		rs = root.Val
	}
	dfsFindBottomLeftValue(root.Left, d+1)
	dfsFindBottomLeftValue(root.Right, d+1)
}
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val
	}
	rightValue := sumOfLeftLeaves(root.Right)
	return leftValue + rightValue
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)
	if leftH == -1 || rightH == -1 {
		return -1
	}
	target := leftH - rightH
	if target > 1 || target < -1 {
		return -1
	}
	return max(leftH, rightH) + 1
}
func isBalanced(root *TreeNode) bool {
	rs := getHeight(root)
	if rs == -1 {
		return false
	}
	return true
}
func binaryTreePaths(root *TreeNode) []string {
	rs := make([]string, 0)
	var travel func(node *TreeNode, path string)
	travel = func(root *TreeNode, path string) {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			rs = append(rs, path)
			return
		}
		if root.Left != nil {
			path += "->"
			travel(root.Left, path)
		}
		if root.Right != nil {
			path += "->"
			travel(root.Right, path)
		}
	}
	travel(root, "")
	return rs
}
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}
	if leftH == rightH {
		return 2<<leftH - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Right != nil && root.Right == nil {
		return minDepth(root.Right) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Right), maxDepth(root.Left)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	return dfsSymmetric(left.Right, right.Left) && dfsSymmetric(left.Left, right.Right)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsSymmetric(root.Right, root.Left)
}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	rs := make([][]int, 0)
	var tmp []int
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[i]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[length:]
		rs = append(rs, tmp)
		tmp = []int{}
	}
	return rs
}

func inorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	var st []*TreeNode
	cur := root
	for len(st) > 0 || cur != nil {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
		} else {
			cur = st[len(st)-1]
			st = st[:len(st)-1]
			rs = append(rs, cur.Val)
			cur = cur.Right
		}
	}
	return rs
}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	if root == nil {
		return rs
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		rs = append(rs, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return rs
}

/*//递归
func preorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		rs = append(rs, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return rs
}
*/
