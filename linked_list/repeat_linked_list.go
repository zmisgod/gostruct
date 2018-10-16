package gostruct

import (
	"fmt"
)

//RepeatLinkedList 单链表
type RepeatLinkedList struct {
	length uint //单链表最大长度
	head   *RListNode
}

//RListNode 每一个节点
type RListNode struct {
	value interface{}
	next  *RListNode
}

//NewRepeatListNode 新建一个NewListNode
func NewRepeatListNode(data interface{}) *RListNode {
	return &RListNode{data, nil}
}

//GetNext ListNode 的下一个节点指针
func (s *RListNode) GetNext() *RListNode {
	return s.next
}

//GetValue 获取ListNode节点的值
func (s *RListNode) GetValue() interface{} {
	if s == nil {
		return "error"
	}
	return s.value
}

//CreateNewRepeatLinkedList 创建一个单链表
func CreateNewRepeatLinkedList() *RepeatLinkedList {
	return &RepeatLinkedList{length: 0, head: NewRepeatListNode(0)}
}

//PreNode 根据ListNode查找前一个节点指针
func (s *RepeatLinkedList) PreNode(n *RListNode) *RListNode {
	if s.head == nil || n == nil {
		return nil
	}
	var i uint
	currentPoint := s.head
	nextPoint := s.head.next
	for ; i < s.length; i++ {
		if nextPoint == n {
			return currentPoint
		}
		currentPoint = nextPoint
		nextPoint = currentPoint.next
	}
	return nil
}

//InsertAfter 在某个节点之后插入
func (s *RepeatLinkedList) InsertAfter(n *RListNode, data interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewRepeatListNode(data)
	oldNext := n.next
	if oldNext == nil {
		oldNext = s.head
	}
	n.next = newNode
	newNode.next = oldNext
	s.length++
	return true
}

//InsertBefore 在n之前插入数据
func (s *RepeatLinkedList) InsertBefore(n *RListNode, data interface{}) bool {
	if n == nil {
		return false
	}
	pre := s.PreNode(n)
	return s.InsertAfter(pre, data)
}

//InsertAfterHead 插入到head后
func (s *RepeatLinkedList) InsertAfterHead(data interface{}) bool {
	return s.InsertAfter(s.head, data)
}

//InsertAfterTail 插入到尾节点后
func (s *RepeatLinkedList) InsertAfterTail(data interface{}) bool {
	curr := s.head
	var i uint
	for ; i < s.length; i++ {
		if curr.next != nil {
			curr = curr.next
		}
	}
	return s.InsertAfter(curr, data)
}

//InsertBeforeTail 插入到尾节点前插入
func (s *RepeatLinkedList) InsertBeforeTail(data interface{}) bool {
	curr := s.head
	var i uint
	for ; i < s.length; i++ {
		if curr.next != nil {
			curr = curr.next
		}
	}
	return s.InsertAfter(curr, data)
}

//FindByIndex 根据索引找到对应的node
func (s *RepeatLinkedList) FindByIndex(index uint) *RListNode {
	//因为index 从0 开始，而length为链表的长度，从 1开始，所以
	//第一个数据的length = 1 index = 0
	//第二个数据的length = 2 index = 1
	//所以需要 <
	if s.length <= index {
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
func (s *RepeatLinkedList) DeleteNode(n *RListNode) bool {
	if n == nil {
		return false
	}
	if s.length <= 0 {
		return false
	}
	pre := s.PreNode(n)
	if pre == nil {
		return false
	}
	pre.next = n.next
	n = nil
	s.length--
	return true
}

//Traverse 遍历
func (s *RepeatLinkedList) Traverse() {
	if s.length > 0 {
		node := s.head.next
		var i uint
		for ; i < s.length; i++ {
			fmt.Println(node)
			node = node.next
		}
	}
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", s.length))
}
