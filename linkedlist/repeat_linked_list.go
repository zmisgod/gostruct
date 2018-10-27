package linkedlist

import (
	"fmt"
)

//解决处理约瑟夫问题
//人们站在一个等待被处决的圈子里。 计数从圆圈中的指定点开始，并沿指定方向围绕圆圈进行。
//在跳过指定数量的人之后，执行下一个人。 对剩下的人重复该过程，从下一个人开始，朝同一方向跳过相同数量的人，直到只剩下一个人，并被释放。
//问题即，给定人数、起点、方向和要跳过的数字，选择初始圆圈中的位置以避免被处决。

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
		return nil
	}
	return s.value
}

//CreateNewRepeatLinkedList 创建一个单链表
func CreateNewRepeatLinkedList() *RepeatLinkedList {
	node := NewRepeatListNode(0)
	node.next = node
	return &RepeatLinkedList{length: 0, head: node}
}

//PreNode 根据ListNode查找前一个节点指针
func (s *RepeatLinkedList) PreNode(n *RListNode) *RListNode {
	if n == nil {
		return nil
	}
	currentPoint := s.head
	nextPoint := s.head.next
	for nextPoint != s.head {
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
	if s.head.next == s.head {
		s.head.next = newNode
	}
	//之前的next指向head
	if n.next == s.head {
		n.next = newNode
		newNode.next = s.head
	} else {
		n.next = newNode
		newNode.next = oldNext
	}
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
	for curr.next != s.head {
		curr = curr.next
	}
	return s.InsertAfter(curr, data)
}

//InsertBeforeTail 插入到尾节点前插入
func (s *RepeatLinkedList) InsertBeforeTail(data interface{}) bool {
	curr := s.head.next
	for curr != s.head {
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
	pre.next = n.next
	n = nil
	s.length--
	return true
}

//Traverse 遍历
func (s *RepeatLinkedList) Traverse() {
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", s.length))
	if s.length > 0 {
		node := s.head.next
		for node != s.head {
			fmt.Println(node)
			node = node.next
		}
	}
}
