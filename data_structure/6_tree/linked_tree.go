package tree

import (
	"container/list"
	"fmt"
)

type LinkedTree struct {
	Value  int         // 值
	Parent *LinkedTree // 父亲节点
	Left   *LinkedTree // 左节点
	Right  *LinkedTree // 右节点
	Level  int         // 层高
}

// 创建树的根节点
func NewLinkedTreeRoot(data int) *LinkedTree {
	return &LinkedTree{
		data,
		nil,
		nil,
		nil,
		0,
	}
}

// 添加
func (tree *LinkedTree) Add(data int) {
	node := &LinkedTree{
		data,
		nil,
		nil,
		nil,
		0,
	}
	pointer := tree
	level := 0
	for pointer != nil {
		if data > pointer.Value {
			if pointer.Right == nil {
				node.Parent = pointer
				node.Level = level
				pointer.Right = node
				break
			}
			pointer = pointer.Right

		} else if data < pointer.Value {
			if pointer.Left == nil {
				node.Parent = pointer
				node.Level = level
				pointer.Left = node
				break
			}

			pointer = pointer.Left
		} else {
			break
		}
	}
}

// 查找
func (tree *LinkedTree) Find(data int) *LinkedTree {
	var node *LinkedTree
	node = nil
	pointer := tree

	for pointer != nil {
		if data > pointer.Value {
			pointer = pointer.Right
		} else if data < pointer.Value {
			pointer = pointer.Left
		} else {
			node = pointer
			break
		}
	}

	return node
}

// 移除
func (tree *LinkedTree) Remove(data int) {
	current := tree.Find(data)
	var parent *LinkedTree
	if current != nil {
		parent = current.Parent

		if parent.Left == current {
			parent.Left = nil
		}

		if parent.Right == current {
			parent.Right = nil
		}
	}

	// 通过层序遍历将后续子节点加入树
	queue := list.New()
	queue.PushBack(current)
	var ele *list.Element
	var node *LinkedTree
	for queue.Front() != nil {
		ele = queue.Front()
		node = ele.Value.(*LinkedTree)

		if node != current {
			tree.Add(node.Value)
		}

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		queue.Remove(ele)
	}
}

// 树高
func (tree *LinkedTree) Height() int {
	return tree.GetNodeHeight(tree)
}

func (tree *LinkedTree) GetNodeHeight(node *LinkedTree) int {
	if node == nil {
		return 0
	}

	leftHeigh := tree.GetNodeHeight(node.Left)
	rightHigh := tree.GetNodeHeight(node.Right)
	if leftHeigh > rightHigh {
		return leftHeigh + 1
	}

	return rightHigh + 1
}

// 前序遍历
func (tree *LinkedTree) PreorderTraversal(node *LinkedTree, position string) {
	if node != nil {
		fmt.Printf("-- %s %d \n", position, node.Value)
		tree.PreorderTraversal(node.Left, "左")
		tree.PreorderTraversal(node.Right, "右")
	}
}

// 中序遍历
func (tree *LinkedTree) MiddleOrderTraversal(node *LinkedTree, position string) {
	if node != nil {
		tree.MiddleOrderTraversal(node.Left, "左")
		fmt.Printf("-- %s %d \n", position, node.Value)
		tree.MiddleOrderTraversal(node.Right, "右")
	}
}

// 后序遍历
func (tree *LinkedTree) PostOrderTraversal(node *LinkedTree, position string) {
	if node != nil {
		tree.PostOrderTraversal(node.Left, "左")
		tree.PostOrderTraversal(node.Right, "右")
		fmt.Printf("%s %d \n", position, node.Value)
	}
}

// 层序遍历
func (tree *LinkedTree) LevelOrderTraversal() {
	queue := list.New()
	queue.PushBack(tree)
	var ele *list.Element
	var node *LinkedTree
	for queue.Front() != nil {
		ele = queue.Front()
		node = ele.Value.(*LinkedTree)
		fmt.Printf("%d \n", node.Value)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		queue.Remove(ele)
	}
}
