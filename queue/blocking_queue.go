package queue

//BlockingQueue 阻塞队列
type BlockingQueue struct {
	length int
	data   chan interface{}
}

//NewBlockingQueue 创建一个阻塞队列
func NewBlockingQueue(maxSize uint) *BlockingQueue {
	return &BlockingQueue{length:0, data: make(chan interface{}, maxSize)}
}

//Enqueue 入队
func (s *BlockingQueue) Enqueue(value interface{}) error {
	select {
	case s.data <- value:
		s.length++
		return nil
	}
}

//Dequeue 出队
func (s *BlockingQueue) Dequeue() (interface{}){
	var data interface{}
	select {
	case data = <-s.data:
		s.length--
		return data
	}
}

func (s *BlockingQueue) isFull() bool{
	if len(s.data) >= s.length{
		return true
	}
	return false
}

func (s *BlockingQueue) isEmpty() bool{
	if s.length == 0{
		return true
	}
	return false
}