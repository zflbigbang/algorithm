package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var f, s *ListNode
	if lenA > lenB {
		step = lenA - lenB
		f, s = headA, headB
	} else {
		step = lenB - lenA
		f, s = headB, headA
	}
	for i := 0; i < step; i++ {
		f = f.Next
	}
	for f != s {
		f = f.Next
		s = s.Next
	}
	return s
}
func detectCycle(head *ListNode) *ListNode {
	s := head
	f := head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if f == s {
			index1 := f
			index2 := head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	s := dummyHead
	f := dummyHead
	n++
	for i := 0; i < n && f != nil; i++ {
		f = f.Next
	}
	for f != nil {
		f = f.Next
		s = s.Next
	}
	s.Next = s.Next.Next
	return dummyHead.Next
}
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp2 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp1
		tmp1.Next = tmp2
		cur = cur.Next.Next
	}
	return dummyHead.Next
}

// 递归写法
func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}
func reverse(pre *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	return reverse(cur, tmp)
}

/* 双指针写法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}*/

type SingleNode struct {
	Val  int
	Next *SingleNode
}
type MyLinkedList struct {
	dummyHead *SingleNode
	size      int
}

func Constructor() MyLinkedList {
	node := &SingleNode{
		0,
		nil,
	}
	return MyLinkedList{
		node,
		0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &SingleNode{Val: val}
	node.Next = this.dummyHead.Next
	this.dummyHead.Next = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	node := &SingleNode{Val: val}
	cur.Next = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index >= this.size {
		return
	}
	node := &SingleNode{Val: val}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.size--
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
