package queue

import (
	"fmt"
	"testing"
)

func TestNormalQueueOne(t *testing.T) {
	queue := NewNormalQueue(4)
	fmt.Println(queue.Enqueue(111))
	fmt.Println(queue.Enqueue(222))
	fmt.Println(queue.Enqueue(333))
	fmt.Println(queue.Enqueue(444))
	fmt.Println(queue.Enqueue(555))

	queue.Traverse()

	fmt.Println("----------")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println("+++++++++")
	fmt.Println(queue.head)
	fmt.Println(queue.length)
	fmt.Println(queue.tail)
}

func TestNormalQueueTwo(t *testing.T) {
	queue := NewNormalQueue(4)
	queue.Enqueue(111)
	queue.Enqueue(222)
	queue.Enqueue(333)
	queue.Enqueue(444)
	queue.Enqueue(555)

	fmt.Println("----------")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println(queue.Enqueue(111))
	fmt.Println(queue.Enqueue(222))
	fmt.Println(queue.Enqueue(333))
	fmt.Println(queue.Enqueue(444))
	fmt.Println(queue.Enqueue(555))

	queue.Traverse()

	fmt.Println("----------")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
