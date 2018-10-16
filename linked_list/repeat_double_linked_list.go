package gostruct

import (
	"fmt"
)

//RepeatDoublueLinkedList 双向链表
type RepeatDoublueLinkedList struct {
	length uint
	head   *RDListNode
}

//RDListNode 双向链表
type RDListNode struct {
	value interface{}
	next  *RDListNode
	pre   *RDListNode
}

//NewRepeatDoubleListNode 创建一个node
func NewRepeatDoubleListNode(value interface{}) *RDListNode {
	return &RDListNode{value, nil, nil}
}

//GetNext 下一个node指针
func (d *RDListNode) GetNext() *RDListNode {
	return d.next
}

//GetPre 上一个node指针
func (d *RDListNode) GetPre() *RDListNode {
	return d.pre
}

//GetValue 此node的值
func (d *RDListNode) GetValue() interface{} {
	if d == nil {
		return "ERROR"
	}
	return d.value
}

//CreateNewRepeatDoublueLinkedList 创建一个双向链表
func CreateNewRepeatDoublueLinkedList() *RepeatDoublueLinkedList {
	return &RepeatDoublueLinkedList{length: 0, head: NewRepeatDoubleListNode(0)}
}

//PreNode 查找上一个node
func (d *RepeatDoublueLinkedList) PreNode(n *RDListNode) *RDListNode {
	return n.pre
}

//NextNode 根据指针查找下一个node
func (d *RepeatDoublueLinkedList) NextNode(n *RDListNode) *RDListNode {
	return n.next
}

//FindByIndex 根据index查找
func (d *RepeatDoublueLinkedList) FindByIndex(index uint) *RDListNode {
	//1 0
	if d.length >= index {
		current := d.head.next
		var i uint
		for ; i < index; i++ {
			current = current.next
		}
		return current
	}
	return nil
}

//InsertAfter 在之后插入
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *RepeatDoublueLinkedList) InsertAfter(n *RDListNode, value interface{}) bool {
	if n == nil || d.head == nil {
		return false
	}
	newNode := NewRepeatDoubleListNode(value)
	nextNode := n.next

	n.next = newNode
	newNode.pre = n
	newNode.next = nextNode
	if nextNode != nil {
		nextNode.pre = newNode
	}
	d.length++
	return true
}

//InsertAfterHead 插入head
func (d *RepeatDoublueLinkedList) InsertAfterHead(value interface{}) bool {
	return d.InsertAfter(d.head, value)
}

//Insert 顺序插入
func (d *RepeatDoublueLinkedList) Insert(value interface{}) bool {
	node := d.head
	for node.next != nil {
		node = node.next
	}
	return d.InsertAfter(node, value)
}

//InsertBefore 在之前插入
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *RepeatDoublueLinkedList) InsertBefore(n *RDListNode, value interface{}) bool {
	node := n.pre
	return d.InsertAfter(node, value)
}

//DeleteNode 删除node
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *RepeatDoublueLinkedList) DeleteNode(n *RDListNode) bool {
	if n == nil {
		return false
	}
	if d.length <= 0 {
		return false
	}
	preNode := n.pre
	nextNode := n.next
	n = nil
	preNode.next = nextNode
	nextNode.pre = nextNode
	d.length--
	return true
}

//Traverse 遍历list
func (d *RepeatDoublueLinkedList) Traverse() {
	if d.length > 0 {
		node := d.head.next
		for node != nil {
			fmt.Println(node.GetValue())
			node = node.next
		}
	}
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", d.length))
}
