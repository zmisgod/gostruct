package binary_tree

import (
	"errors"
	"fmt"
)

type BTree struct {
	root   *TreeNode
	length uint
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value uint
}

func (node *TreeNode) GetMinNode(parent *TreeNode, dir uint8) (*TreeNode, uint8, *TreeNode) {
	if node == nil {
		return nil, 0, parent
	}
	if node.left == nil && node.right == nil {
		return node, dir, parent
	}
	if node.left != nil {
		parent = node
		return node.left.GetMinNode(parent, 1)
	}
	return node, 2, parent
}

func (node *TreeNode) GetMaxNode(parent *TreeNode, dir uint8) (*TreeNode, uint8, *TreeNode) {
	if node == nil {
		return nil, 0, parent
	}
	if node.left == nil && node.right == nil {
		return node, dir, parent
	}
	if node.right != nil {
		parent = node
		return node.right.GetMinNode(parent, 2)
	}
	return nil, 0, parent
}

func (node *TreeNode) InOrder() {
	if node == nil {
		return
	}
	node.left.InOrder()
	fmt.Printf("%v \n", node.value)
	node.right.InOrder()
}

func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	fmt.Printf("%v \n", node.value)
	node.left.PreOrder()
	node.right.PreOrder()
}

func (node *TreeNode) PostOrder() {
	if node == nil {
		return
	}
	node.left.PostOrder()
	node.right.PostOrder()
	fmt.Printf("%v \n", node.value)
}

func NewNode(value uint) *TreeNode {
	return &TreeNode{value: value}
}

func NewBTree() *BTree {
	return &BTree{}
}

func (tree *BTree) Add(value uint) *TreeNode {
	//当前搜索到的节点
	node := tree.root
	//父节点
	var parent *TreeNode
	for node != nil {
		parent = node
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	nowNode := NewNode(value)
	if parent == nil {
		tree.root = nowNode
	} else if value < parent.value {
		parent.left = nowNode
	} else {
		parent.right = nowNode
	}
	tree.length += 1
	return parent
}

func (node *TreeNode) FindValue(value uint, parent *TreeNode) (*TreeNode, uint8, *TreeNode) {
	var dir uint8
	//父节点
	for node != nil {
		if node.value > value {
			parent = node
			dir = 1
			return node.left.FindValue(value, parent)
		} else if node.value < value {
			parent = node
			dir = 2
			return node.right.FindValue(value, parent)
		} else {
			return node, dir, parent
		}
	}
	return nil, dir, parent
}

func (tree *BTree) FindValue(value uint) (*TreeNode, uint8, *TreeNode) {
	node := tree.root
	parent := node
	return node.FindValue(value, parent)
}

func (tree *BTree) Delete(value uint) error {
	searchNode, _, parentNode := tree.FindValue(value)
	if searchNode == nil {
		return errors.New("DO NOT FIND THIS VALUE")
	}
	if searchNode.left != nil && searchNode.right != nil {
		minNode := searchNode.right
		minParent := searchNode
		for minNode.left != nil {
			minParent = minNode
			minNode = minNode.left
		}
		searchNode.value = minNode.value
		searchNode = minNode
		parentNode = minParent
	}
	//需要删除节点的子节点
	var sonSearchNode *TreeNode
	if searchNode.left != nil {
		sonSearchNode = searchNode.left
	} else if searchNode.right != nil {
		sonSearchNode = searchNode.right
	}else{
		sonSearchNode = nil
	}

	if parentNode == nil {
		fmt.Println("----111")
		//删除的就是根节点
		tree.root = sonSearchNode
	} else if parentNode.left == searchNode {
		parentNode.left = sonSearchNode
	} else {
		parentNode.right = sonSearchNode
	}
	return nil
}

//中序遍历
func (tree *BTree) InOrderTraverse() {
	if tree.root == nil {
		return
	}
	tree.root.InOrder()
}

//先序遍历
func (tree *BTree) PreOrderTraverse() {
	if tree.root == nil {
		return
	}
	tree.root.PreOrder()
}

//后序遍历
func (tree *BTree) PostOrderTraverse() {
	if tree.root == nil {
		return
	}
	tree.root.PostOrder()
}

//获取叶子节点
func (tree *BTree) GetLeavesPoint() uint {
	return 0
}

//是否为满二叉树
func (tree *BTree) IsFullBTree() bool {
	return false
}

//是否为完全二叉树
func (tree *BTree) IsCompletedBTree() bool {
	return false
}

//获取此树的深度
func (tree *BTree) GetDeep(index uint) uint {

	return 0
}

//获取树的深度
func (tree *BTree) GetTreeDepth() uint {
	return 0
}

//获取树的层
func (tree *BTree) GetTreeLevel() uint {
	return 0
}

//获取父节点
func (tree *BTree) GetFatherPoint() uint {
	return 0
}

//获取子节点
func (tree *BTree) GetSonPoint() uint {
	return 0
}

//获取兄弟节点
func (tree *BTree) GetBrotherPoint(node *TreeNode) uint {
	return 0
}

//获取树的高度
func (tree *BTree) GetTreeHeight() uint {
	return 0
}
