package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queueCap := 2
	myqueue := CreateQueue(queueCap)
	var queone = "this is a test"
	myqueue.EnQueue(queone)
	var quetwo = "this is a test 2"
	myqueue.EnQueue(quetwo)
	myqueue.Traverse()

	var queueDequeue interface{}
	myqueue.DeQueue(&queueDequeue)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeue)
	fmt.Println("_+_+_+_+")

	lenght := myqueue.Llength()
	fmt.Println(lenght)

	var queueDequeueR interface{}
	myqueue.DeQueue(&queueDequeueR)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeueR)
	fmt.Println("_+_+_+_+")
}
