package gostruct

import (
	"fmt"
)

//SingleLinkedList 单链表
type SingleLinkedList struct {
	length uint //单链表最大长度
	head   *SListNode
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
func CreateNewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{length: 0, head: NewSingleListNode(0)}
}

//InsertAfter 在某个节点之后插入
func (s *SingleLinkedList) InsertAfter(n *SListNode, data interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewSingleListNode(data)
	oldNext := n.next
	n.next = newNode
	newNode.next = oldNext
	s.length++
	return true
}

//PreNode 根据ListNode查找前一个节点指针
func (s *SingleLinkedList) PreNode(n *SListNode) *SListNode {
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

//InsertBefore 在n之前插入数据
func (s *SingleLinkedList) InsertBefore(n *SListNode, data interface{}) bool {
	if n == nil || s.head == nil {
		return false
	}
	if n == nil {
		return false
	}
	pre := s.PreNode(n)
	return s.InsertAfter(pre, data)
}

//InsertAfterHead 插入到head后
func (s *SingleLinkedList) InsertAfterHead(data interface{}) bool {
	return s.InsertAfter(s.head, data)
}

//InsertAfterTail 插入到尾节点后
func (s *SingleLinkedList) InsertAfterTail(data interface{}) bool {
	curr := s.head
	for curr.next != nil {
		curr = curr.next
	}
	return s.InsertAfter(curr, data)
}

//InsertBeforeTail 插入到尾节点前插入
func (s *SingleLinkedList) InsertBeforeTail(data interface{}) bool {
	curr := s.head
	for curr.next != nil {
		curr = curr.next
	}
	return s.InsertAfter(curr, data)
}

//FindByIndex 根据索引找到对应的node
func (s *SingleLinkedList) FindByIndex(index uint) *SListNode {
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
func (s *SingleLinkedList) DeleteNode(n *SListNode) bool {
	if n == nil {
		return false
	}
	if s.length <= 0 {
		return false
	}
	pre := s.PreNode(n)
	pre.next = n.next
	n = nil
	s.length--
	return true
}

//Traverse 遍历
func (s *SingleLinkedList) Traverse() {
	if s.length > 0 {
		node := s.head.next
		for node != nil {
			fmt.Println(node)
			node = node.next
		}
	}
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", s.length))
}
