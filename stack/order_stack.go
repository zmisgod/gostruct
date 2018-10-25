package gostruct

//OrderStack 顺序栈
type OrderStack struct {
	data   []interface{}
	length uint
}

//NewOrderStack 创建一个顺序栈
func NewOrderStack(maxStackSize int) *OrderStack {
	return &OrderStack{data: make([]interface{}, maxStackSize, maxStackSize), length: 0}
}

//Push 插入一个数据
func (s *OrderStack) Push(value interface{}) bool {
	if s.IsFull() {
		return false
	}
	s.data[s.length] = value
	s.length++
	return true
}

//Pop 从栈尾处取一条数据
func (s *OrderStack) Pop() interface{} {
	if s.IsEmpty() {
		return false
	}
	s.length--
	data := s.data[s.length]
	s.data[s.length] = nil

	return data
}

//olen 判断数组的长度
func (s *OrderStack) olen() int {
	return len(s.data)
}

//IsFull 判断数组是否满了
func (s *OrderStack) IsFull() bool {
	if int(s.length) >= s.olen() {
		return true
	}
	return false
}

//IsEmpty 判断数组是否为空
func (s *OrderStack) IsEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}
