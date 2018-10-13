package gostruct

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stackSize := 2
	mystack := CreateStack(stackSize)
	var stackone = "zmisgod test 1"
	mystack.Push(stackone)
	var stacktwo = "zmisgod test 2"
	mystack.Push(stacktwo)

	var popStack interface{}
	mystack.Pop(&popStack)
	fmt.Println("before pop")
	fmt.Println(popStack)
	fmt.Println("after pop")
	mystack.StackTraverse()
}
