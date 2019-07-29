package queue

import (
	"fmt"
	"testing"
)

func TestBlockingQueue(t *testing.T) {
	queue := NewBlockingQueue(4)
	fmt.Printf("out %v \n", queue.Dequeue())
	queue.Enqueue(222)
	fmt.Printf("out %v \n", queue.Dequeue())
	queue.Enqueue(333)
	queue.Enqueue(444)
	queue.Enqueue(555)
	fmt.Printf("out %v \n", queue.Dequeue())
	fmt.Printf("out %v \n", queue.Dequeue())
	fmt.Printf("out %v \n", queue.Dequeue())
	queue.Enqueue(666)
	queue.Enqueue(777)
	queue.Enqueue(888)
	queue.Enqueue(999)
	queue.Enqueue(1000)
	queue.Enqueue(1111)
	fmt.Printf("out %v \n", queue.Dequeue())
	fmt.Printf("out %v \n", queue.Dequeue())
	fmt.Printf("out %v \n", queue.Dequeue())
	fmt.Printf("out %v \n", queue.Dequeue())
}
