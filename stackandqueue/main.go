package main

func main() {

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

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

/*func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}*/

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
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

func (this *MyQueue) Peek() int {
	value := this.Pop()
	if value == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, value)
	return value
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

// MyQueue 泛型实现队列
/*type MyQueue[T interface{}] struct {
	elements []T
}

func Constructor() MyQueue[int] {
	return MyQueue[int]{}
}

func (q *MyQueue[T]) put(value T) {
	q.elements = append(q.elements, value)
}

func (q *MyQueue[T]) Pop() T {
	var value T
	if len(q.elements) > 0 {
		value = q.elements[0]
		q.elements = q.elements[1:]
		return value
	}
	return value
}

func (q *MyQueue[T]) Peek() T {
	var value T
	if len(q.elements) > 0 {
		value = q.elements[0]
		return value
	}
	return value
}

func (q *MyQueue[T]) Empty() bool {
	if len(q.elements) == 0 {
		return true
	}
	return false
}
*/
