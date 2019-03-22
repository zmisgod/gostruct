package linkedlist

import (
	"errors"
	"fmt"
)

//Array 数组
type Array struct {
	length uint
	data   []int
}

//CreateNewArray 创建一个数组
func CreateNewArray(maxLenght uint) *Array {
	if maxLenght <= 0 {
		return nil
	}
	return &Array{0, make([]int, maxLenght, maxLenght)}
}

//Len 数组长度
func (a *Array) Len() uint {
	return a.length
}

//Insert 插入数据
func (a *Array) Insert(index uint, value int) bool {
	//检查数组容量是否超过大小
	if a.Len() >= uint(cap(a.data)) {
		return false
	}
	//检查数组是否越界
	// 0 1
	if index >= a.length && a.IsIndexOutOfRange(index) {
		return false
	}
	//如果插入index，意味着之前的index后的数据都需要向后移位
	// 2=>1,3=>2,4 => 3,5 => 4
	// 插入3 => 5,
	//2=> 1, 3 => 5, 4 => 2, 5 =>3 , 6 => 4
	for i := index; i < a.Len(); i++ {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = value
	a.length++
	return true
}

//IsIndexOutOfRange 检查数组是否越界
func (a *Array) IsIndexOutOfRange(index uint) bool {
	//1 0
	//2 1
	if a.length >= index {
		return false
	}
	return true
}

//FindByIndex 根据index查value
func (a *Array) FindByIndex(n uint) (int, error) {
	//1 0
	//2 1
	if !a.IsIndexOutOfRange(n) {
		return a.data[n], nil
	}
	return 0, errors.New("is out of range")
}

//InsertToTail 插入到尾部
func (a *Array) InsertToTail(value int) bool {
	if a.Len() >= uint(cap(a.data)) {
		return false
	}
	return a.Insert(a.Len(), value)
}

//DeleteByIndex 根据索引删除
func (a *Array) DeleteByIndex(index uint) bool {
	if a.IsIndexOutOfRange(index) {
		return false
	}
	var i uint
	//只需要将index后面的数据向前移动一位即可
	// 0 1
	// 1 2
	// 2 3
	for ; i < a.Len()-1; i++ {
		if i >= index {
			a.data[i] = a.data[i+1]
		}
	}
	//a.data[a.Len() -1]
	a.length--
	return true
}

//Traverse 遍历
func (a *Array) Traverse() {
	if a.length > 0 {
		var i uint
		for ; i < a.Len(); i++ {
			fmt.Println(a.data[i])
		}
	}
	fmt.Println(fmt.Sprintf("single linked list's length is %d ", a.length))
}
