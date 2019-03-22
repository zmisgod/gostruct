package queue

import (
	"flag"
	"fmt"
	"testing"
)

func TestBlockingQueue(t *testing.T) {
	//queue := NewBlockingQueue(4)
	//queue.Enqueue(222)
	//queue.Enqueue(333)
	//queue.Enqueue(444)
	//queue.Enqueue(555)
	//fmt.Println(queue.Dequeue())
	//fmt.Println(queue.Dequeue())
	//fmt.Println(queue.Dequeue())
	//fmt.Println(queue.Dequeue())
	//close(queue.data)
	var input int
	flag.IntVar(&input, "input", 1, "test")
	var chan1 chan int
	var stop int
	for stop == 0 {
		data := <- chan1
		fmt.Println(data)
	}
}