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
	node := RDListNode{value, nil, nil}
	node.pre = &node
	node.next = &node
	return &node
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
		return nil
	}
	return d.value
}

//CreateNewRepeatDoublueLinkedList 创建一个双向链表
func CreateNewRepeatDoublueLinkedList() *RepeatDoublueLinkedList {
	node := NewRepeatDoubleListNode(0)
	node.next = node
	node.pre = node
	return &RepeatDoublueLinkedList{length: 0, head: node}
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
		current := d.head.GetNext()
		var i uint
		for ; i < index; i++ {
			current = current.GetNext()
		}
		return current
	}
	return nil
}

//InsertAfter 在之后插入
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *RepeatDoublueLinkedList) InsertAfter(n *RDListNode, value interface{}) bool {
	if n == nil {
		return false
	}
	newNode := NewRepeatDoubleListNode(value)
	nextNode := n.GetNext()
	n.next = newNode
	newNode.pre = n
	newNode.next = nextNode
	nextNode.pre = newNode
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
	var i uint
	for ; i < d.length; i++ {
		node = node.GetNext()
	}
	return d.InsertAfter(node, value)
}

//InsertBefore 在之前插入
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *RepeatDoublueLinkedList) InsertBefore(n *RDListNode, value interface{}) bool {
	node := n.GetPre()
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
	preNode := n.GetPre()
	nextNode := n.GetNext()
	n = nil
	preNode.next = nextNode
	nextNode.pre = preNode
	d.length--
	return true
}

//Traverse 遍历list
func (d *RepeatDoublueLinkedList) Traverse() {
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", d.length))
	if d.length > 0 {
		node := d.head.GetNext()
		var i uint
		for ; i < d.length; i++ {
			fmt.Println(node.GetValue())
			node = node.GetNext()
		}
	}
}
