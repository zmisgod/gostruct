package gostruct

import (
	"fmt"
	"testing"
)

func TestAfterHead(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.InsertAfterHead("111")
	singleList.InsertAfterHead("222")
	singleList.InsertAfterHead("333")
	singleList.Traverse()
}

func TestInsertToTail(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	singleList.Traverse()
}

func TestFindByIndex(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(3)
	fmt.Println(node.GetValue())
}

func TestInsertAfterNode(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(1)
	singleList.InsertAfter(node, "444")
	singleList.Traverse()
}

func TestInsertBeforeNode(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(1)
	singleList.InsertBefore(node, "444")
	singleList.Traverse()
}

func TestDelete(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.InsertAfterTail("111")
	singleList.InsertAfterTail("222")
	singleList.InsertAfterTail("333")
	node := singleList.FindByIndex(12)
	singleList.DeleteNode(node)
	singleList.Traverse()
}
