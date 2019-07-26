package binary_tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewBTree()
	_ = tree.Add(2)
	_ = tree.Add(3)
	_ = tree.Add(4)
	_ = tree.Add(5)
	_ = tree.Add(1)
	_ = tree.Add(7)
	_ = tree.Add(8)
	fmt.Println(tree.Delete(8))
	tree.PreOrderTraverse()
	//fmt.Println("----")
	//tree.InOrderTraverse()
	//fmt.Println("----")
	//tree.PostOrderTraverse()
	//fmt.Println("----")
}
