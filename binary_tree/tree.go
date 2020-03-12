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

func (node *TreeNode) LeafPrint() {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		fmt.Printf("%v \n", node.value)
	}
	node.left.LeafPrint()
	node.right.LeafPrint()
}

func (node *TreeNode) CheckFather(count *int) interface{} {
	if node.left != nil || node.right != nil {
		*count++
		if node.left != nil {
			node.left.CheckFather(count)
		}
		if node.right != nil {
			node.right.CheckFather(count)
		}
	}
	return nil
}

//是否为满二叉树
func (node *TreeNode) IsFullTree() bool {
	isFull := false

	return isFull
}

//是否为完全二叉树
func (node *TreeNode) IsComTree() bool {
	return false
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

func (node *TreeNode) FindValue(value uint, parent *TreeNode) (*TreeNode, *TreeNode) {
	//父节点
	for node != nil {
		if node.value > value {
			parent = node
			return node.left.FindValue(value, parent)
		} else if node.value < value {
			parent = node
			return node.right.FindValue(value, parent)
		} else {
			return node, parent
		}
	}
	return nil, parent
}

func (tree *BTree) FindValue(value uint) (*TreeNode, *TreeNode) {
	node := tree.root
	parent := node
	return node.FindValue(value, parent)
}

func (tree *BTree) Delete(value uint) error {
	searchNode, parentNode := tree.FindValue(value)
	if searchNode == nil {
		return errors.New("DO NOT FIND THIS VALUE")
	}
	//当节点的左右子树都有节点的情况
	if searchNode.left != nil && searchNode.right != nil {
		minNode := searchNode.right
		minParent := searchNode
		for minNode.left != nil {
			minParent = minNode
			minNode = minNode.left
		}
		//将需要删除的节点的value的值替换此节点右边最小的节点的值
		searchNode.value = minNode.value
		//然后此问题就变成了删除最左节点（最小）的问题
		searchNode = minNode
		parentNode = minParent
	}

	//当节点的左右子树只有一个或没有节点的情况
	var sonSearchNode *TreeNode
	if searchNode.left != nil {
		sonSearchNode = searchNode.left
	} else if searchNode.right != nil {
		sonSearchNode = searchNode.right
	} else {
		sonSearchNode = nil
	}

	if parentNode == nil {
		//如果没有父节点，说明删除的就是根节点
		tree.root = sonSearchNode
	} else if parentNode.left == searchNode {
		//如果父节点的左子树等于删除的节点，则父节点的左子树为当前的子节点
		parentNode.left = sonSearchNode
	} else {
		//否则为右子树的子节点
		parentNode.right = sonSearchNode
	}
	return nil
}

//打印中序遍历
func (tree *BTree) InOrderTraverse() {
	if tree.root == nil {
		return
	}
	tree.root.InOrder()
}

//打印先序遍历
func (tree *BTree) PreOrderTraverse() {
	if tree.root == nil {
		return
	}
	tree.root.PreOrder()
}

//打印后序遍历
func (tree *BTree) PostOrderTraverse() {
	if tree.root == nil {
		return
	}
	tree.root.PostOrder()
}

//打印叶子节点（没有子节点）
func (tree *BTree) GetLeavesPoint() {
	tree.root.LeafPrint()
}

//获取父节点
func (tree *BTree) GetFatherPoint() int {
	var i int
	tree.root.CheckFather(&i)
	return i
}

//是否为满二叉树（所有层的节点数都达到最大）
func (tree *BTree) IsFullBTree() bool {
	return false
}

//是否为完全二叉树（倒数第二层节点都达到最大）
func (tree *BTree) IsCompletedBTree() bool {
	return false
}

//获取树的深度
func (tree *BTree) GetTreeDepth() uint {
	return 0
}

//获取树的层
func (tree *BTree) GetTreeLevel() uint {
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
