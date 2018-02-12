# goqueue

simple queue useing golang

## install

```
go get github.com/zmisgod/goqueue
```

for test 

```
package main 

import (
	"fmt"
	"github.com/zmisgod/goqueue"
)

func main () {
	queueCap := 2
	myqueue := goqueue.CreateQueue(queueCap)
	var queone goqueue.Queue
	queone.Name = "zmisgo"
	queone.Job = "php"
	queone.Age = 12
	myqueue.EnQueue(queone)
	var quetwo goqueue.Queue
	quetwo.Name = "zmisgod"
	quetwo.Job = "php"
	quetwo.Age = 13
	myqueue.EnQueue(quetwo)
	myqueue.QueueTraverse()

	var queueDequeue goqueue.Queue
	myqueue.DeQueue(&queueDequeue)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeue)
	fmt.Println("_+_+_+_+")

	lenght := myqueue.QueueLength()
	fmt.Println(lenght)

	var queueDequeueR goqueue.Queue
	myqueue.DeQueue(&queueDequeueR)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeueR)
	fmt.Println("_+_+_+_+")
}
```

## Postscript

<a href="https://zmis.me/detail_1378">数据结构探险—队列golang篇</a>