package linkedlist

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	lru.Put("1", "1")
	lru.Put("2", "2")
	lru.Put("3", "3")
	one := lru.Get("1")
	two := lru.Get("2")
	three := lru.Get("3")
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
}