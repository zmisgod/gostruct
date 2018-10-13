// @环形队列
package gostruct

import (
	"fmt"
)

//Queue 队列
type Queue struct {
	QueuePtr []*interface{}
	Length   int
	Capacity int
	Heap     int
	Tail     int
}

//CreateQueue 创建队列
func CreateQueue(capacity int) *Queue {
	return &Queue{Length: 0, Capacity: capacity, Heap: 0, Tail: 0}
}

//Clear 清空队列
func (q *Queue) Clear() {
	q.Heap = 0
	q.Tail = 0
	q.Length = 0
}

//Empty 判断是否为空队列
func (q *Queue) Empty() bool {
	if q.Length == 0 {
		return true
	}
	return false
}

//Llength 队列的长度
func (q *Queue) Llength() int {
	return q.Length
}

//Full 判断当前队列是否已满
func (q *Queue) Full() bool {
	if q.Length == q.Capacity {
		return true
	}
	return false
}

//EnQueue 入队
func (q *Queue) EnQueue(queue interface{}) bool {
	//当队列满了就没法入队
	if q.Full() {
		return false
	}
	//将这个入队的地址附加到QueuePtr的数组指针中
	q.QueuePtr = append(q.QueuePtr, &queue)
	//尾部向前加一
	q.Tail++
	//这边可能出现数组越界(尾部的索引越界)，所以取尾部与容量的取余操作
	q.Tail = q.Tail % q.Capacity
	//队列的长度加一
	q.Length++
	return true
}

//DeQueue 出队 queue这里为指针，这样才可以返回相应的出队的Queue对象
func (q *Queue) DeQueue(queue *interface{}) bool {
	//判断队列是否为空
	if q.Empty() {
		return false
	}
	//将传过来的地址的值改写为获取到的值
	*queue = *q.QueuePtr[q.Heap]
	//队列头的索引加一
	q.Heap++
	//和出队一样，这边可能出现数组越界(头部的索引越界)，所以取头部与容量的取余操作
	q.Heap = q.Heap % q.Capacity
	//队列的长度减一
	q.Length--
	return true
}

//Traverse 遍历队列
func (q *Queue) Traverse() {
	//遍历队列的起始是队列的数据解构保存的Heap（头索引）字段
	//遍历的长度是当前队列的长度+头索引 （为了保证每次循环的次数相同）
	//当头索引为0，队列容量为4， 队列长度为4时，遍历的个数为 4 + 0
	//当头索引为1，队列容量为4， 队列长度为3时，遍历的个数为 3 + 1
	//当头索引为2，队列容量为4， 队列长度为2时，遍历的个数为 2 + 2
	//当头索引为3，队列容量为4， 队列长度为1时，遍历的个数为 1 + 3
	for i := q.Heap; i < q.Length+q.Heap; i++ {
		res := q.QueuePtr[i%q.Capacity]
		fmt.Println(*res)
	}
}
