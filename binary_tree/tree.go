package binary_tree

import (
	"errors"
	"fmt"
	"math"
)

type BTree struct {
	root   *TreeNode
	length uint
}

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	value  uint
	parent *TreeNode
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

func NewNode(value uint, parent *TreeNode) *TreeNode {
	return &TreeNode{value: value, parent: parent}
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
	nowNode := NewNode(value, parent)
	if parent == nil {
		tree.root = nowNode
	}else if value < parent.value {
		parent.left = nowNode
	} else {
		parent.right = nowNode
	}
	tree.length += 1
	return parent
}

func (tree *BTree) FindValue(value uint) *TreeNode {
	node := tree.root
	for node != nil {
		if value > node.value {
			node = node.right
		} else if node.value == value {
			return node
		} else {
			node = node.left
		}
	}
	return nil
}

func (tree *BTree) Delete(value uint) error {
	searchNode := tree.FindValue(value)
	if searchNode == nil {
		return errors.New("没有找到该数字")
	}
	isLeft := false
	//情况1 这个节点没有子节点
	if searchNode.left == nil && searchNode.right == nil {
		searchNode = nil
		if searchNode.parent.left == searchNode {
			isLeft = true
		}
		if isLeft {
			searchNode.parent.left = nil
		}else{
			searchNode.parent.right = nil
		}
	}
	//情况1 这个节点有且仅有一个子节点
	if (searchNode.left != nil && searchNode.right == nil) || (searchNode.left == nil && searchNode.right != nil) {
		//仅存在左节点
		if searchNode.left != nil && searchNode.right == nil {
			if searchNode.parent != nil {
				searchNode.parent.right = searchNode.right
			}
			searchNode = nil
			tree.length -= 1
			return nil
		}
		//仅存在右节点
		if searchNode.left != nil && searchNode.right == nil {
			if searchNode.parent != nil {
				searchNode.parent.left = searchNode.left
			}
			searchNode = nil
			tree.length -= 1
			return nil
		}
	}

	//情况3 这个节点左子树与右子树都有子节点
	if searchNode.left != nil && searchNode.right != nil {
		temp := searchNode
		saveTemp :=searchNode
		if searchNode.left != nil {
			searchNode := searchNode.left
			searchNode.value = value
			temp.left = searchNode.left
			for ;searchNode.left != nil ;searchNode= searchNode.left  {

			}
			searchNode.left = saveTemp
			saveTemp.left = nil
		}else{
			saveTemp.value = searchNode.value
			temp.left = searchNode.left
		}
	}
	return nil
}

//获取树的高度
func (tree *BTree) GetTreeHeight() uint {
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

func (tree *BTree) GetDeep(index uint) uint {
	deep := uint(math.Ceil(math.Log2(float64(index + 1))))
	return deep
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
