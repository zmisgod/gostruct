package queue

import "fmt"

//CircularQueue 环形队列
type CircularQueue struct {
	data   []interface{}
	length uint
	heap   uint
	tail   uint
}

//CreateQueue 创建队列
func CreateCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{length: 0, data: make([]interface{}, capacity, capacity), heap: 0, tail: 0}
}

func (s *CircularQueue) Enqueue(value interface{}) bool {
	if s.isFull() {
		return false
	}
	s.data[s.tail] = value
	tail := s.tail
	tail++
	s.tail = tail % uint(s.olen())
	s.length++
	return false
}

func (s *CircularQueue) Dequeue() interface{} {
	if s.isEmpty() {
		return nil
	}
	data := s.data[s.heap]
	s.data[s.heap] = nil
	heap := s.heap
	heap++
	s.heap = heap % uint(s.olen())
	s.length--
	return data
}

func (s *CircularQueue) Traverse() {
	if s.length > 0 {
		fmt.Println(s.data)
	}
	fmt.Println(fmt.Sprintf(" circular queue length is %d ", s.length))
}

func (s *CircularQueue) len() int {
	return int(s.length)
}

func (s *CircularQueue) isFull() bool {
	if (s.tail+1)%uint(s.olen()) == s.heap {
		return true
	}
	return false
}

func (s *CircularQueue) isEmpty() bool {
	if s.heap == s.tail {
		return true
	}
	return false
}

func (s *CircularQueue) olen() int {
	return len(s.data)
}
