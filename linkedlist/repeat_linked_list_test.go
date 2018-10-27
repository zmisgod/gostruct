package linkedlist

import (
	"fmt"
	"testing"
)

func TestRAfterHead(t *testing.T) {
	singleList := CreateNewRepeatLinkedList()
	singleList.InsertAfterHead("111")
	singleList.InsertAfterHead("222")
	singleList.InsertAfterHead("333")
	singleList.Traverse()
}

func TestRInsertToTail(t *testing.T) {
	singleList := CreateNewRepeatLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	singleList.Traverse()
}

func TestRFindByIndex(t *testing.T) {
	singleList := CreateNewRepeatLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	singleList.Traverse()
	node := singleList.FindByIndex(0)
	fmt.Println(node.GetValue())
}

func TestRInsertAfterNode(t *testing.T) {
	singleList := CreateNewRepeatLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(1)
	singleList.InsertAfter(node, "444")
	singleList.Traverse()
}

func TestRInsertBeforeNode(t *testing.T) {
	singleList := CreateNewRepeatLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(1)
	singleList.InsertBefore(node, "444")
	singleList.Traverse()
}

func TestRDelete(t *testing.T) {
	singleList := CreateNewRepeatLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(0)
	singleList.DeleteNode(node)
	node1 := singleList.FindByIndex(1)
	singleList.DeleteNode(node1)
	// node2 := singleList.FindByIndex(0)
	// singleList.DeleteNode(node2)
	singleList.Traverse()
}
