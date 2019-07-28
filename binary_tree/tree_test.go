package binary_tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewBTree()
	_ = tree.Add(4)
	_ = tree.Add(3)
	_ = tree.Add(5)
	_ = tree.Add(6)
	_ = tree.Add(1)
	_ = tree.Add(0)
	_ = tree.Add(2)
	_ = tree.Add(7)
	_ = tree.Add(8)
	if err := tree.Delete(3); err != nil {
		fmt.Println(err)
	}
	if err := tree.Delete(4); err != nil {
		fmt.Println(err)
	}
	fmt.Println("----")
	tree.PreOrderTraverse()
	//fmt.Println("----")
	//tree.InOrderTraverse()
	//fmt.Println("----")
	//tree.PostOrderTraverse()
	//fmt.Println("----")
}
