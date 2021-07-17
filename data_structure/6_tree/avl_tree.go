package tree

import (
	"errors"
)

type AVLTree struct {
	Value int
	Depth int
	Left  *AVLTree
	Right *AVLTree
}

// 搜索
func (tree *AVLTree) Search(data int) *AVLTree {
	pointer := tree
	for pointer != nil {
		if data > pointer.Value {
			pointer = pointer.Left
		} else if data < pointer.Value {
			pointer = pointer.Right
		} else {
			break
		}
	}

	return pointer
}

// 最小值
func (tree *AVLTree) Min() (min int, err error) {
	if tree == nil {
		return 0, errors.New("空树，不存在最小值")
	}

	pointer := tree
	for pointer.Left != nil {
		pointer = pointer.Left
	}

	return pointer.Value, nil
}

// 最大值
func (tree *AVLTree) Max() (max int, err error) {
	if tree == nil {
		return 0, errors.New("空树 不存在最大值")
	}

	pointer := tree
	for pointer.Right != nil {
		pointer = pointer.Right
	}

	return pointer.Value, nil
}

// 获取节点层高
func (node *AVLTree) GetDepth() int {
	if node == nil {
		return 0
	}

	return node.Depth
}

// 左旋
func (node *AVLTree) LeftRotate() *AVLTree {
	grandson := node.Right
	node.Left = grandson.Left
	grandson.Left = node

	node.Depth = max(node.Left.GetDepth(), node.Right.GetDepth()) + 1
	grandson.Depth = max(grandson.Left.GetDepth(), grandson.Right.GetDepth()) + 1

	return grandson
}

// 右旋
func (node *AVLTree) RightRotate() *AVLTree {
	grandson := node.Left
	node.Right = grandson.Right
	grandson.Right = node

	node.Depth = max(node.Left.GetDepth(), node.Right.GetDepth()) + 1
	grandson.Depth = max(grandson.Left.GetDepth(), grandson.Right.GetDepth()) + 1

	return grandson
}

// 右旋再左旋
func (node *AVLTree) RightThenLeftRotate() *AVLTree {
	next := node.Right.RightRotate()
	node.Right = next

	return node.LeftRotate()
}

// 左旋再右旋
func (node *AVLTree) LeftThenRightRotate() *AVLTree {
	next := node.Left.LeftRotate()
	node.Left = next

	return node.RightRotate()
}

// 重新平衡
func (node *AVLTree) Rebalance() *AVLTree {
	var (
		left  int = node.Left.GetDepth()
		right int = node.Right.GetDepth()
	)

	if right-left == 2 {
		if node.Right.Right.GetDepth() > node.Right.Left.GetDepth() {
			node = node.LeftRotate()
		} else {
			node = node.RightThenLeftRotate()
		}
	} else if left-right == 2 {
		if node.Left.Left.GetDepth() > node.Left.Right.GetDepth() {
			node = node.RightRotate()
		} else {
			node = node.RightThenLeftRotate()
		}
	}

	return node
}

// 添加
func (tree *AVLTree) Insert(data int) *AVLTree {
	if tree == nil {
		return &AVLTree{
			data,
			1,
			nil,
			nil,
		}
	}

	if data > tree.Value {
		tree.Right = tree.Right.Insert(data)
		return tree.Rebalance()
	}

	if data < tree.Value {
		tree.Left = tree.Left.Insert(data)
		return tree.Rebalance()
	}

	return tree
}

// 删除
func (tree *AVLTree) Delete(data int) *AVLTree {
	if tree == nil {
		return tree
	}

	if data > tree.Value {
		tree.Right = tree.Right.Delete(data)
	} else if data < tree.Value {
		tree.Left = tree.Left.Delete(data)
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Right.Value = tree.Right.GetMin()
			tree.Right = tree.Right.Delete(tree.Value)
		} else if tree.Left != nil {
			tree = tree.Left
		} else {
			tree = tree.Right
		}
	}

	if tree != nil {
		tree.Depth = max(tree.Left.GetDepth(), tree.Right.GetDepth()) + 1
		tree = tree.Rebalance()
	}

	return tree
}

// 子树最小值
func (tree *AVLTree) GetMin() int {
	if tree.Left == nil {
		return tree.Value
	}

	return tree.Left.GetMin()
}

// 前序遍历
func (tree *AVLTree) PreorderTraversal() []int {
	data := make([]int, 0)
	var traversal func(node *AVLTree)
	traversal = func(node *AVLTree) {
		if node == nil {
			return
		}

		data = append(data, node.Value)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(tree)
	return data
}

// 两数比较求最大值
func max(first int, second int) int {
	if first > second {
		return first
	}

	return second
}
