package linkedlist

import (
	"fmt"
)

// 使用双向循环链表
// Put的时间复杂度为o(1)
// Get的时间复杂度为o(n)
type LRURepeatDoubleLinkedList struct {
	maxCap uint
	nowCap uint
	head *LRUCacheNode
}

type LRUCacheNode struct {
	key string
	value interface{}
	pre *LRUCacheNode
	next *LRUCacheNode
}

func NewLRUCache(maxCap uint) *LRURepeatDoubleLinkedList{
	node := newNode("", 0)
	node.pre = node
	node.next = node
	return &LRURepeatDoubleLinkedList{maxCap:maxCap, head:node}
}

func newNode(key string , value interface{})*LRUCacheNode {
	return &LRUCacheNode{key:key, value:value}
}

func (n *LRUCacheNode) PreValue() (string, interface{}){
	return n.pre.key, n.pre.value
}

func (n *LRUCacheNode) NextValue() (string, interface{}) {
	return n.next.key, n.next.value
}

func (cache *LRURepeatDoubleLinkedList) Put(key string ,value interface{}) {
	if key == "" {
		return
	}
	node := newNode(key, value)
	if cache.maxCap <= cache.nowCap {
		last := cache.head.pre
		last.next = nil
		cache.head.pre = last.pre
		last.pre.next = cache.head
		last.pre = nil
	}

	first := cache.head.next
	first.pre = node
	node.next = first
	cache.head.next = node
	node.pre = cache.head
	if cache.maxCap > cache.nowCap {
		cache.nowCap += 1
	}
}

func (cache *LRURepeatDoubleLinkedList)  Get(key string ) interface{} {
	if cache.nowCap > 0 {
		var i uint
		now := cache.head.next
		for ; i < cache.nowCap; i++ {
			if now.key == key {
				//需要将此key设置到头部
				now.pre.next = now.next
				now.next.pre = now.pre

				//将now设置在cache头部
				first := cache.head.next
				first.pre = now
				now.next = first
				cache.head.next = now
				now.pre = cache.head

				return now.value
			}else{
				now = now.next
			}
		}
	}
	return -1
}

func (cache *LRURepeatDoubleLinkedList) Traverse() {
	if cache.nowCap > 0 {
		var i uint
		now := cache.head.next
		for ; i < cache.nowCap; i++ {
			fmt.Printf("key = %s value %v\n", now.key, now.value)
			now = now.next
		}
	}
}
