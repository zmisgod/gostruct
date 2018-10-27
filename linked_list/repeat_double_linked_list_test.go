package linked_list

import (
	"fmt"
	"testing"
)

func TestRInsertAfterNodeDouble(t *testing.T) {
	doubleList := CreateNewRepeatDoublueLinkedList()
	doubleList.InsertAfterHead("111")
	doubleList.InsertAfterHead("222")
	doubleList.InsertAfterHead("333")
	doubleList.Traverse()
}

func TestRInsertDouble(t *testing.T) {
	doubleList := CreateNewRepeatDoublueLinkedList()
	doubleList.Insert("111")
	doubleList.Insert("222")
	doubleList.Insert("333")
	doubleList.Insert("444")
	doubleList.Insert("555")
	node := doubleList.FindByIndex(21)
	fmt.Println(node.GetValue())
	fmt.Println("++======")
	doubleList.Traverse()
}

func TestRInsertAfterDouble(t *testing.T) {
	doubleList := CreateNewRepeatDoublueLinkedList()
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

func TestRInsertBeforeDouble(t *testing.T) {
	doubleList := CreateNewRepeatDoublueLinkedList()
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

func TestRDeleteNodeDouble(t *testing.T) {
	doubleList := CreateNewRepeatDoublueLinkedList()
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
	node = doubleList.FindByIndex(0)
	doubleList.DeleteNode(node)
	doubleList.Traverse()
}
