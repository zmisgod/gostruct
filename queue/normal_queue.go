package queue

import "fmt"
//NormalQueue 普通队列
type NormalQueue struct {
	length uint //当前队列长度
	tail   uint //尾指针
	head   uint //头指针
	data   []interface{}
}
//NewNormalQueue 创建一个普通队列
func NewNormalQueue(maxSize uint) *NormalQueue {
	return &NormalQueue{length: 0, data: make([]interface{}, maxSize, maxSize), tail: 0, head: 0}
}

//Enqueue 入队
func (s *NormalQueue) Enqueue(value interface{}) bool {
	if s.olen() == int(s.tail){
		if s.head == 0 {
			return false
		}else{
			tail := s.tail - s.head
			for i:=s.head; i< s.tail;i++  {
				s.data[s.tail - i] = s.data[i]
			}
			s.head = 0
			s.tail = tail
		}
	}
	s.data[s.tail] = value
	s.tail++
	s.length++
	return true
}

//Dequeue 出队
func (s *NormalQueue) Dequeue() interface{} {
	if s.IsEmpty() {
		return nil
	}
	data := s.data[s.head]
	s.data[s.head] = nil
	s.head++
	s.length--
	return data
}

//Traverse 遍历
func (s *NormalQueue) Traverse() {
	if s.length > 0 {
		for i := s.head; i < s.tail; i++ {
			fmt.Println(s.data[i])
		}
	}
	fmt.Println(fmt.Sprintf(" single linked list's length is %d ", s.length))
}

//olen 判断数组的长度
func (s *NormalQueue) olen() int {
	return len(s.data)
}

//IsFull 判断队列是否满了
func (s *NormalQueue) IsFull() bool {
	if s.olen() == int(s.tail) {
		return true
	}
	return false
}

//IsEmpty 判断队列是否为空
func (s *NormalQueue) IsEmpty() bool {
	if s.tail == s.head {
		return true
	}
	return false
}
