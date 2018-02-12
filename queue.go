package goqueue 

import (
	"fmt"
)

/**
* 环形队列
*/

//MyQueue 队列
type MyQueue struct {
	QueuePtr []*Queue
	Length int
	Capacity int
	Heap int
	Tail int
}

//Queue 队列信息
type Queue struct {
	Name string
	Age int
	Job string
}

//CreateQueue 创建队列
func CreateQueue(capacity int) *MyQueue {
	return &MyQueue{Length:0, Capacity:capacity, Heap:0, Tail:0}
}

//DestoryQueue 销毁队列 golang不需要内存回收
// func (q *MyQueue) DestoryQueue() {
// }

//ClearQueue 清空队列
func (q *MyQueue) ClearQueue() {
	q.Heap = 0
	q.Tail = 0
	q.Length = 0
}

//QueueEmpty 判断是否为空队列
func (q *MyQueue) QueueEmpty() bool {
	if q.Length == 0 {
		return true
	}
	return false
}

//QueueLength 队列的长度
func (q *MyQueue) QueueLength() int{
	return q.Length
}

//QueueFull 判断当前队列是否已满
func (q *MyQueue) QueueFull() bool {
	if q.Length == q.Capacity {
		return true
	}
	return false
}

//EnQueue 入队
func (q *MyQueue) EnQueue(queue Queue) bool {
	//当队列满了就没法入队
	if q.QueueFull() {
		return false
	}
	//将这个入队的地址附加到QueuePtr的数组指针中
	q.QueuePtr = append(q.QueuePtr, &queue)
	//尾部向前加一
	q.Tail ++
	//这边可能出现数组越界(尾部的索引越界)，所以取尾部与容量的取余操作
	q.Tail = q.Tail % q.Capacity
	//队列的长度加一
	q.Length ++
	return true
}

//DeQueue 出队 queue这里为指针，这样才可以返回相应的出队的Queue对象
func (q * MyQueue) DeQueue(queue *Queue) bool {
	//判断队列是否为空
	if q.QueueEmpty() {
		return false
	}
	//将传过来的地址的值改写为获取到的值
	*queue = *q.QueuePtr[q.Heap]
	//队列头的索引加一
	q.Heap ++
	//和出队一样，这边可能出现数组越界(头部的索引越界)，所以取头部与容量的取余操作
	q.Heap = q.Heap % q.Capacity
	//队列的长度减一
	q.Length --
	return true
}

//QueueTraverse 遍历队列
func (q *MyQueue) QueueTraverse() {
	//遍历队列的起始是队列的数据解构保存的Heap（头索引）字段
	//遍历的长度是当前队列的长度+头索引 （为了保证每次循环的次数相同）
	//当头索引为0，队列容量为4， 队列长度为4时，遍历的个数为 4 + 0
	//当头索引为1，队列容量为4， 队列长度为3时，遍历的个数为 3 + 1
	//当头索引为2，队列容量为4， 队列长度为2时，遍历的个数为 2 + 2
	//当头索引为3，队列容量为4， 队列长度为1时，遍历的个数为 1 + 3
	for i := q.Heap; i < q.Length + q.Heap; i++ {
		res := q.QueuePtr[i % q.Capacity]
		res.ShowQueue()
	}
}

//ShowQueue 显示队列信息
func (res *Queue) ShowQueue() {
	fmt.Println(res.Age)
	fmt.Println(res.Name)
	fmt.Println(res.Job)
}