package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	queue := CreateCircularQueue(4)
	queue.Enqueue(111)
	queue.Enqueue(222)
	queue.Enqueue(333)
	queue.Enqueue(444)
	queue.Enqueue(555)
	queue.Enqueue(666)
	queue.Traverse()
	fmt.Println(queue.Dequeue())

	queue.Enqueue(444)
	queue.Enqueue(555)
	queue.Enqueue(666)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	queue.Enqueue(555)
	queue.Enqueue(666)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
