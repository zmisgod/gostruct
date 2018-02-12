package gostruct

import (
	"fmt"
	"testing"
)

func TestQueue (t *testing.T) {
	queueCap := 2
	myqueue := CreateQueue(queueCap)
	var queone Queue
	queone.Name = "zmisgod"
	queone.Job = "queue test1"
	queone.Age = 12
	myqueue.EnQueue(queone)
	var quetwo Queue
	quetwo.Name = "zmisgod"
	quetwo.Job = "queue test2"
	quetwo.Age = 13
	myqueue.EnQueue(quetwo)
	myqueue.QueueTraverse()

	var queueDequeue Queue
	myqueue.DeQueue(&queueDequeue)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeue)
	fmt.Println("_+_+_+_+")

	lenght := myqueue.QueueLength()
	fmt.Println(lenght)

	var queueDequeueR Queue
	myqueue.DeQueue(&queueDequeueR)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeueR)
	fmt.Println("_+_+_+_+")
}