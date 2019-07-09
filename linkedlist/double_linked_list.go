package linkedlist

import (
	"fmt"
)

//DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	length uint
	head   *DListNode
}

//DListNode 双向链表
type DListNode struct {
	value interface{}
	next  *DListNode
	pre   *DListNode
}

//NewDoubleListNode 创建一个node
func NewDoubleListNode(value interface{}) *DListNode {
	return &DListNode{value, nil, nil}
}

//GetNext 下一个node指针
func (d *DListNode) GetNext() *DListNode {
	return d.next
}

//GetPre 上一个node指针
func (d *DListNode) GetPre() *DListNode {
	return d.pre
}

//GetValue 此node的值
func (d *DListNode) GetValue() interface{} {
	if d == nil {
		return nil
	}
	return d.value
}

//CreateNewDoubleLinkedList 创建一个双向链表
func CreateNewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{length: 0, head: NewDoubleListNode(0)}
}

//PreNode 查找上一个node
func (d *DoubleLinkedList) PreNode(n *DListNode) *DListNode {
	return n.pre
}

//NextNode 根据指针查找下一个node
func (d *DoubleLinkedList) NextNode(n *DListNode) *DListNode {
	return n.next
}

//FindByIndex 根据index查找
func (d *DoubleLinkedList) FindByIndex(index uint) *DListNode {
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
func (d *DoubleLinkedList) InsertAfter(n *DListNode, value interface{}) bool {
	if n == nil || d.head == nil {
		return false
	}
	newNode := NewDoubleListNode(value)
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
func (d *DoubleLinkedList) InsertAfterHead(value interface{}) bool {
	return d.InsertAfter(d.head, value)
}

//Insert 顺序插入
func (d *DoubleLinkedList) Insert(value interface{}) bool {
	node := d.head
	for node.next != nil {
		node = node.next
	}
	return d.InsertAfter(node, value)
}

//InsertBefore 在之前插入
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *DoubleLinkedList) InsertBefore(n *DListNode, value interface{}) bool {
	node := n.pre
	return d.InsertAfter(node, value)
}

//DeleteNode 删除node
//              pre5 pre6
// [pre1 pre2] [nil  nil ] [pre3 pre4]
func (d *DoubleLinkedList) DeleteNode(n *DListNode) bool {
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
	nextNode.pre = preNode
	d.length--
	return true
}

//Traverse 遍历list
func (d *DoubleLinkedList) Traverse() {
	if d.length > 0 {
		node := d.head.next
		for node != nil {
			fmt.Println(node.GetValue())
			node = node.next
		}
	}
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", d.length))
}
