package gostruct

import (
	"fmt"
)

//MyStack stack
type MyStack struct {
	StackPtr []*Stack //栈空间
	Size int //容量
	Top int //栈顶，栈中元素的个数
}

//Stack 插入栈的解构
type Stack struct {
	Name string
	Age int
	Job string
}

//CreateStack 创建Stack
func CreateStack(size int) *MyStack {
	return &MyStack{Size: size, Top:0, StackPtr: make([]*Stack, size)}
}

// func (s *MyStack) DestoryStack() {
// }

//StackEmpty 判断栈是否为空
func (s *MyStack) StackEmpty() bool {
	if s.Top == 0 {
		return true
	}
	return false
}

//StackFull 判断栈是否满了
func (s *MyStack) StackFull() bool {
	if s.Top >= s.Size  {
		return true
	}
	return false
}

//ClearStack 清空栈
func (s *MyStack) ClearStack() {
	s.Top = 0
}

//StackLength 栈的长度
func (s *MyStack) StackLength() int{
	return s.Top
}

//Push 入栈
func (s *MyStack) Push(ele Stack) bool {
	if s.StackFull() {
		return false
	}
	//先赋值再加Top
	s.StackPtr[s.Top] = &ele
	s.Top ++
	return true
}

//Pop 出栈
func (s *MyStack) Pop(ele *Stack) bool {
	if s.StackEmpty() {
		return false
	}
	//先减Top再取值
	s.Top --
	*ele = *s.StackPtr[s.Top]
	return true
}

//StackTraverse 遍历栈
func (s *MyStack) StackTraverse() {
	for i:=0; i< s.Top; i++ {
		s.StackPtr[i].ShowQueue()
	}
}

//ShowQueue 显示队列信息
func (res *Stack) ShowQueue() {
	fmt.Println(res.Age)
	fmt.Println(res.Name)
	fmt.Println(res.Job)
}