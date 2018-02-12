# gostruct

- [x] simple queue
- [ ] simple stack

## install

```
go get github.com/zmisgod/gostruct
```

for test

```
package main

import (
	"fmt"
	"github.com/zmisgod/gostruct"
)

func main () {
	queueCap := 2
	myqueue := gostruct.CreateQueue(queueCap)
	var queone gostruct.Queue
	queone.Name = "zmisgo"
	queone.Job = "php"
	queone.Age = 12
	myqueue.EnQueue(queone)
	var quetwo gostruct.Queue
	quetwo.Name = "zmisgod"
	quetwo.Job = "php"
	quetwo.Age = 13
	myqueue.EnQueue(quetwo)
	myqueue.QueueTraverse()

	var queueDequeue gostruct.Queue
	myqueue.DeQueue(&queueDequeue)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeue)
	fmt.Println("_+_+_+_+")

	lenght := myqueue.QueueLength()
	fmt.Println(lenght)

	var queueDequeueR gostruct.Queue
	myqueue.DeQueue(&queueDequeueR)
	fmt.Println("_+_+_+_+")
	fmt.Println(queueDequeueR)
	fmt.Println("_+_+_+_+")
}
```

## Postscript

<a href="https://zmis.me/detail_1378">数据结构探险—队列golang篇</a>