package linkedlist

import (
	"testing"
)

func TestAArrayPush(t *testing.T) {
	array := CreateNewArray(10)
	array.Insert(0, 111)
	array.Insert(1, 222)
	array.Insert(2, 333)
	array.Insert(3, 444)
	array.Traverse()
	array.DeleteByIndex(0)
	array.DeleteByIndex(1)
	array.Traverse()

	// array.InsertToTail(777)
	// array.InsertToTail(888)
	// array.InsertToTail(999)
	// array.Traverse()
}
