package gostruct

import (
	"fmt"
	"os"
	"testing"
)

func TestSInsertNode(t *testing.T) {
	staticList, err := NewStaticLinkedList(6)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		staticList.Insert(11111)
		staticList.Insert(22222)
		staticList.Insert(33333)
		staticList.Traverse()
	}
}
