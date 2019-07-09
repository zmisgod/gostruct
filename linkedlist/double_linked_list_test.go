package linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertAfterNodeDouble(t *testing.T) {
	doubleList := CreateNewDoubleLinkedList()
	doubleList.InsertAfterHead("111")
	doubleList.InsertAfterHead("222")
	doubleList.InsertAfterHead("333")
	doubleList.Traverse()
}
func TestInsertDouble(t *testing.T) {
	doubleList := CreateNewDoubleLinkedList()
	doubleList.Insert("111")
	doubleList.Insert("222")
	doubleList.Insert("333")
	doubleList.Insert("444")
	doubleList.Insert("555")
	node := doubleList.FindByIndex(5)
	fmt.Println(node.GetValue())
	doubleList.Traverse()
}

func TestInsertAfterDouble(t *testing.T) {
	doubleList := CreateNewDoubleLinkedList()
	doubleList.Insert("111")
	doubleList.Insert("222")
	doubleList.Insert("333")
	doubleList.Insert("444")
	doubleList.Insert("555")
	node := doubleList.FindByIndex(2)
	doubleList.InsertAfter(node, 666)
	doubleList.InsertAfter(node, 777)
	doubleList.Traverse()
}

func TestInsertBeforeDouble(t *testing.T) {
	doubleList := CreateNewDoubleLinkedList()
	doubleList.Insert("111")
	doubleList.Insert("222")
	doubleList.Insert("333")
	doubleList.Insert("444")
	doubleList.Insert("555")
	node := doubleList.FindByIndex(2)
	doubleList.InsertBefore(node, 666)
	doubleList.InsertBefore(node, 777)
	doubleList.Traverse()
}

func TestDeleteNodeDouble(t *testing.T) {
	doubleList := CreateNewDoubleLinkedList()
	doubleList.Insert("111")
	doubleList.Insert("222")
	doubleList.Insert("333")
	doubleList.Insert("444")
	doubleList.Insert("555")
	node := doubleList.FindByIndex(0)
	doubleList.DeleteNode(node)
	node = doubleList.FindByIndex(0)
	doubleList.DeleteNode(node)
	node = doubleList.FindByIndex(0)
	doubleList.DeleteNode(node)
	doubleList.Traverse()
}
