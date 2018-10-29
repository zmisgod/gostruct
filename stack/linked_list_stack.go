package stack

import (
	"fmt"
)

//SingleLinkedList 单链表
type LinkedListStack struct {
	top  uint //栈顶
	head *SListNode
}

//SListNode 每一个节点
type SListNode struct {
	value interface{}
	next  *SListNode
}

//NewSingleListNode 新建一个NewListNode
func NewSingleListNode(data interface{}) *SListNode {
	return &SListNode{data, nil}
}

//GetNext ListNode 的下一个节点指针
func (s *SListNode) GetNext() *SListNode {
	return s.next
}

//GetValue 获取ListNode节点的值
func (s *SListNode) GetValue() interface{} {
	if s == nil {
		return nil
	}
	return s.value
}

//CreateNewSingleLinkedList 创建一个单链表
func CreateNewSingleLinkedList() *LinkedListStack {
	return &LinkedListStack{top: 0, head: NewSingleListNode(0)}
}

//InsertAfter 在某个节点之后插入
func (s *LinkedListStack) insertAfter(n *SListNode, data interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewSingleListNode(data)
	oldNext := n.next
	n.next = newNode
	newNode.next = oldNext
	s.top++
	return true
}

//PreNode 根据ListNode查找前一个节点指针
func (s *LinkedListStack) preNode(n *SListNode) *SListNode {
	if s.head == nil || n == nil {
		return nil
	}
	var i uint
	currentPoint := s.head
	nextPoint := s.head.next
	for ; i < s.top; i++ {
		if nextPoint == n {
			return currentPoint
		}
		currentPoint = nextPoint
		nextPoint = currentPoint.next
	}
	return nil
}

//InsertBeforeTail 插入到尾节点前插入
func (s *LinkedListStack) Push(data interface{}) bool {
	curr := s.head
	for curr.next != nil {
		curr = curr.next
	}
	return s.insertAfter(curr, data)
}

//Pop 删除
func (s *LinkedListStack) Pop() *SListNode {
	length := s.top
	length --
	node := s.findByIndex(length)
	s.deleteNode(node)
	return node
}

//FindByIndex 根据索引找到对应的node
func (s *LinkedListStack) findByIndex(index uint) *SListNode {
	//因为index 从0 开始，而length为链表的长度，从 1开始，所以
	//第一个数据的length = 1 index = 0
	//第二个数据的length = 2 index = 1
	//所以需要 <
	if s.top <= index {
		return nil
	}
	var i uint
	curr := s.head.next
	for ; i < index; i++ {
		curr = curr.next
	}
	return curr
}

//DeleteNode 删除node
func (s *LinkedListStack) deleteNode(n *SListNode) bool {
	if n == nil {
		return false
	}
	if s.top <= 0 {
		return false
	}
	pre := s.preNode(n)
	pre.next = n.next
	n = nil
	s.top--
	return true
}

//Traverse 遍历
func (s *LinkedListStack) Traverse() {
	if s.top > 0 {
		node := s.head.next
		for node != nil {
			fmt.Println(node)
			node = node.next
		}
	}
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", s.top))
}
