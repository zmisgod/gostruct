package stack

import (
	"fmt"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	singleList := CreateNewSingleLinkedList()
	singleList.Push("111")
	singleList.Push("222")
	singleList.Push("333")
	singleList.Traverse()
	fmt.Println(singleList.Pop().GetValue())
	fmt.Println(singleList.Pop().GetValue())
	fmt.Println(singleList.Pop().GetValue())
	fmt.Println(singleList.Pop().GetValue())
	fmt.Println(singleList.Pop().GetValue())
}
