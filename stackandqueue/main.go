package main

import (
	"container/heap"
	"strconv"
)

func main() {

}
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	rs := make([]int, k)

	for _, num := range nums {
		m[num]++
	}
	h := &IHeap{}
	for key, value := range m {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := k - 1; i >= 0; i-- {
		rs[i] = heap.Pop(h).([2]int)[0]
	}
	return rs
}

// 构建小顶堆 sort.Interface
//
//	Push(x any) // add x as element Len()
//	Pop() any
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// 封装单调队列的方式解题
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) front() int {
	return m.queue[0]
}

func (m *MyQueue) back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) pop(val int) {
	if len(m.queue) > 0 && m.front() == val {
		m.queue = m.queue[1:]
	}
}
func (m *MyQueue) push(val int) {
	for i := len(m.queue) - 1; i >= 0; i-- {
		if val > m.queue[i] {
			m.queue = m.queue[:len(m.queue)-1]
		} else {
			break
		}
	}
	m.queue = append(m.queue, val)
}
func maxSlidingWindow(nums []int, k int) []int {
	q := NewMyQueue()
	rs := make([]int, 0)
	for i := 0; i < k; i++ {
		q.push(nums[i])
	}
	rs = append(rs, q.front())
	for i := k; i < len(nums); i++ {
		q.pop(nums[i-k])
		q.push(nums[i])
		rs = append(rs, q.front())
	}
	return rs
}
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err != nil {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num2+num1)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num2*num1)
			case "/":
				stack = append(stack, num2/num1)
			}
		} else {
			stack = append(stack, val)
		}
	}
	return stack[0]
}
func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
func isValid(s string) bool {
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < len(this.queue)-1; i++ {
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/*type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (this MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this MyQueue) Pop() int {
	inLen, outLen := len(this.stackIn), len(this.stackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}
		outLen = len(this.stackOut)
	}
	value := this.stackOut[outLen-1]
	this.stackOut = this.stackOut[:outLen-1]
	return value
}

func (this MyQueue) Peek() int {
	value := this.Pop()
	if value == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, value)
	return value
}

func (this MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
*/

// MyQueue 泛型实现队列
/*type MyQueue[T interface{}] struct {
	elements []T
}

func Constructor() MyQueue[int] {
	return MyQueue[int]{}
}

func (q MyQueue[T]) put(value T) {
	q.elements = append(q.elements, value)
}

func (q MyQueue[T]) Pop() T {
	var value T
	if len(q.elements) > 0 {
		value = q.elements[0]
		q.elements = q.elements[1:]
		return value
	}
	return value
}

func (q MyQueue[T]) Peek() T {
	var value T
	if len(q.elements) > 0 {
		value = q.elements[0]
		return value
	}
	return value
}

func (q MyQueue[T]) Empty() bool {
	if len(q.elements) == 0 {
		return true
	}
	return false
}
*/
