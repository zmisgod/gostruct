package queue

//BlockingQueue 阻塞队列
type BlockingQueue struct {
	length int
	cap    int
	data   chan interface{}
}

//NewBlockingQueue 创建一个阻塞队列
func NewBlockingQueue(maxSize int) *BlockingQueue {
	queue := &BlockingQueue{
		length: 0,
		cap:    maxSize,
		data:   make(chan interface{}, maxSize),
	}
	return queue
}

//Enqueue 入队
func (s *BlockingQueue) Enqueue(value interface{}) {
	if !s.isFull() {
		s.data <- value
		s.length++
	}
}

//Dequeue 出队
func (s *BlockingQueue) Dequeue() interface{} {
	if !s.isEmpty() {
		for v := range s.data {
			s.length--
			return v
		}
	}
	return nil
}

func (s *BlockingQueue) isFull() bool {
	if s.length >= s.cap {
		return true
	}
	return false
}

func (s *BlockingQueue) isEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}
