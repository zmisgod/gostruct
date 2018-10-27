package stack

import (
	"fmt"
	"testing"
)

func TestOrderStack(t *testing.T) {
	os := NewOrderStack(2)
	fmt.Println(os.Push(111))
	fmt.Println(os.Push(222))
	fmt.Println(os.Push(333))
	fmt.Println(os.Pop())
	fmt.Println(os.Pop())
	fmt.Println(os.Pop())
}
