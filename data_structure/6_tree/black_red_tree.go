package tree

import (
	"errors"
	"fmt"
)

type BlackRedTreeNode struct {
	Value  int               // 值
	Color  string            // 颜色
	Parent *BlackRedTreeNode // 父节点
	Left   *BlackRedTreeNode // 左节点
	Right  *BlackRedTreeNode // 右节点
}

type BlackTree struct {
	root *BlackRedTreeNode
}

// 判断是否是空树
func (tree *BlackTree) IsEmpty() bool {
	return tree.root == nil
}

// 添加
func (tree *BlackTree) Insert(data int) {
	node := new(BlackRedTreeNode)
	node.Value = data
	node.Color = "red"

	// 空树则创建根节点
	if tree.IsEmpty() {
		node.Color = "black"
		tree.root = node
		fmt.Printf("插入后前序遍历 %v\n", tree.PreorderTraversal())
		return
	}

	// 插入
	pointer := tree.root
	for pointer != nil {
		if data > pointer.Value {
			if pointer.Right == nil {
				node.Parent = pointer
				pointer.Right = node
				break
			}
			pointer = pointer.Right
		} else if data < pointer.Value {
			if pointer.Left == nil {
				node.Parent = pointer
				pointer.Left = node
				break
			}
			pointer = pointer.Left
		} else {
			break
		}
	}

	// 重新平衡
	tree.InsertRebalance(node)

	fmt.Printf("插入后前序遍历 %v\n", tree.PreorderTraversal())
}

// 插入重平衡
func (tree *BlackTree) InsertRebalance(node *BlackRedTreeNode) {
	var uncle *BlackRedTreeNode
	var parent *BlackRedTreeNode

	// 只有当前节点和父节点都是红色才需要重新平衡
	for node.Color == "red" && node.Parent != nil && node.Parent.Color == "red" {
		parent = node.Parent
		// 获得叔叔节点
		if parent == parent.Parent.Left {
			uncle = parent.Parent.Right
		} else {
			uncle = parent.Parent.Left
		}

		if uncle != nil && uncle.Color == "red" {
			// 如果父节点和叔叔节点都是红色，证明树还处于平衡，
			// 但是不满足上下节点同时是红色，所以需要向上重新渲染
			parent.Color, uncle.Color = "balck", "black"
			node = parent.Parent
			if node == tree.root || node == nil {
				return
			}
			node.Color = "red"
		} else {
			// 如果父节点为红色 叔叔节点为黑色或者nil 证明树已不是完全二叉树， 需要重新平衡

			if parent == parent.Parent.Left {
				// 父节点是祖父节点的左子节点 左边层高搞 则右旋

				if node == parent.Right {
					// 如果插入节点是右孩子 对其父节点先左旋
					node = node.Parent
					tree.LeftRotate(node)
				}

				node = node.Parent
				node.Color = "black"
				node = node.Parent
				node.Color = "red"

				tree.RightRotate(node)
			} else {
				// 父节点是祖父节点的右孩子 右边层高搞 则左旋

				if node == parent.Left {
					// 如果插入节点是左孩子 对其父节点先右旋转
					node = node.Parent
					tree.RightRotate(node)
				}

				node = node.Parent
				node.Color = "black"
				node = node.Parent
				node.Color = "red"

				tree.LeftRotate(node)
			}
		}
	}
}

// 删除
func (tree *BlackTree) Delete(data int) {
	deleteNode := tree.Search(data)
	if deleteNode == nil {
		return
	}

	var remove func(node *BlackRedTreeNode)
	remove = func(node *BlackRedTreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			// 叶子节点之上 节点可以直接删除
			if node == tree.root { // 根
				tree.root = nil
				return
			}

			if node == node.Parent.Left {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		} else if node.Left != nil || node.Right != nil {
			// 只有一个子节点
			var child *BlackRedTreeNode
			if node.Left != nil {
				child = node.Left
			} else {
				child = node.Right
			}

			if node == tree.root { // 根
				child.Parent = nil
				tree.root = child
				return
			}

			child.Parent = node.Parent
			if node == node.Parent.Left {
				node.Parent.Left = child
			} else {
				node.Parent.Right = child
			}
		} else {
			// 同时有左右节点 在右节点找最小值替换给当前节点
			replaceNode := tree.GetReplaceNode(node)
			node.Value = replaceNode.Value
			node.Color = replaceNode.Color
			remove(replaceNode)
		}
	}

	parent := deleteNode.Parent
	var reviseNode *BlackRedTreeNode
	if deleteNode.Parent == nil && (deleteNode.Left != nil || deleteNode.Right != nil) {
		reviseNode = tree.root
	} else if deleteNode != nil && deleteNode == parent.Left {
		reviseNode = parent.Left
	} else if deleteNode != nil && deleteNode == parent.Right {
		reviseNode = parent.Right
	}

	remove(deleteNode)

	// 重新配平
	if deleteNode.Color == "black" {
		tree.DeleteRebalance(reviseNode)
	}

	fmt.Printf("删除后前序遍历 %v\n", tree.PreorderTraversal())
}

// 删除后再平衡
func (tree *BlackTree) DeleteRebalance(node *BlackRedTreeNode) {
	var brother *BlackRedTreeNode
	for node != tree.root && node.Color == "black" {
		if node == node.Parent.Left && node.Parent.Left != nil {
			brother = node.Parent.Right
			if brother.Color == "red" {
				brother.Color = "black"
				node.Parent.Color = "red"
				tree.LeftRotate(node.Parent)
				return
			}

			if brother.Left != nil && brother.Left.Color == "black" &&
				brother.Right != nil && brother.Right.Color == "black" {
				brother.Color = "red"
				node = node.Parent
				return
			}

			if brother.Left != nil && brother.Left.Color == "red" &&
				brother.Right != nil && brother.Right.Color == "black" {
				brother.Color = "red"
				brother.Left.Color = "black"
				tree.RightRotate(brother)
				return
			}

			if brother.Right != nil && brother.Right.Color == "red" {
				brother.Color = "red"
				brother.Right.Color = "black"
				brother.Parent.Color = "black"
				tree.LeftRotate(brother.Parent)
				node = tree.root
				return
			}

		} else if node == node.Parent.Right && node.Parent.Right != nil {
			brother = node.Parent.Left
			if brother.Color == "red" {
				brother.Color = "black"
				node.Parent.Color = "red"
				tree.RightRotate(node.Parent)
				return
			}

			if brother.Left != nil && brother.Left.Color == "black" &&
				brother.Right != nil && brother.Right.Color == "black" {
				brother.Color = "red"
				node = node.Parent
				return
			}

			if brother.Left != nil && brother.Left.Color == "red" &&
				brother.Right != nil && brother.Right.Color == "black" {
				brother.Color = "red"
				brother.Right.Color = "black"
				tree.LeftRotate(brother)
				return
			}

			if brother.Left != nil && brother.Left.Color == "red" {
				brother.Color = "red"
				brother.Left.Color = "black"
				brother.Parent.Color = "black"
				tree.RightRotate(brother.Parent)
				node = tree.root
				return
			}
		} else {
			return
		}
	}
}

// 节点左旋
func (tree *BlackTree) LeftRotate(node *BlackRedTreeNode) {
	// 如果旋转节点没有右孩子 则无法左旋
	if node.Right == nil {
		return
	}

	right := node.Right
	node.Right = right.Left
	if node.Right != nil {
		node.Right.Parent = node
	}

	right.Parent = node.Parent
	if node.Parent == nil {
		tree.root = node
	} else {
		if node.Parent.Left == node {
			node.Parent.Left = right
		} else {
			node.Parent.Right = node
		}
	}

	right.Left = node
	node.Parent = right
}

// 节点左旋
func (tree *BlackTree) RightRotate(node *BlackRedTreeNode) {
	// 如果旋转节点没有左孩子 则无法右旋
	if node.Left == nil {
		return
	}

	left := node.Left
	node.Left = left.Right
	if node.Left != nil {
		node.Left.Parent = node
	}

	left.Parent = node.Parent
	if node.Parent == nil {
		tree.root = node
	} else {
		if node.Parent.Left == node {
			node.Parent.Left = left
		} else {
			node.Parent.Right = node
		}
	}

	left.Right = node
	node.Parent = left
}

// 查找节点
func (tree *BlackTree) Search(data int) *BlackRedTreeNode {
	pointer := tree.root
	for pointer != nil {
		if data > pointer.Value {
			pointer = pointer.Right
		} else if data < pointer.Value {
			pointer = pointer.Left
		} else {
			break
		}
	}

	return pointer
}

// 获取删除可替换节点
func (tree *BlackTree) GetReplaceNode(node *BlackRedTreeNode) (replace *BlackRedTreeNode) {
	if node == nil || node.Parent == nil {
		return
	}

	if node.Right != nil {
		pointer := node.Right.Left
		for pointer != nil {
			pointer = pointer.Left
		}
		return pointer
	}

	for {
		if node == node.Parent.Left {
			return node.Parent
		}
		node = node.Parent
	}
}

// 获取根节点
func (tree *BlackTree) GetRoot() *BlackRedTreeNode {
	if tree == nil {
		return nil
	}

	return tree.root
}

// 获取树高
func (tree *BlackTree) GetHeight() int {
	var height func(node *BlackRedTreeNode) int
	height = func(node *BlackRedTreeNode) int {
		if node == nil {
			return 0
		}

		if node.Left == nil && node.Right == nil {
			return 1
		}

		var (
			left  int = height(node.Left)
			right int = height(node.Right)
		)

		if left > right {
			return left + 1
		}

		return right + 1
	}

	return height(tree.root)
}

// 最小值
func (tree *BlackTree) Min() (min int, err error) {
	if tree == nil || tree.root == nil {
		return 0, errors.New("当前树是空树，不存在最小值")
	}

	pointer := tree.root
	minValue := pointer.Value
	for pointer != nil {
		if pointer.Value < min {
			minValue = pointer.Value
		}
		pointer = pointer.Left
	}

	return minValue, nil
}

//  最大值
func (tree *BlackTree) Max() (max int, err error) {
	if tree == nil || tree.root == nil {
		return 0, errors.New("当前树是空树，不存在最大值")
	}

	pointer := tree.root
	maxValue := pointer.Value
	for pointer != nil {
		if pointer.Value > max {
			maxValue = pointer.Value
		}
		pointer = pointer.Right
	}

	return maxValue, nil
}

// 前序遍历
func (tree *BlackTree) PreorderTraversal() []int {
	data := make([]int, 0)
	var traversal func(node *BlackRedTreeNode)
	traversal = func(node *BlackRedTreeNode) {
		if node == nil {
			return
		}

		data = append(data, node.Value)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(tree.root)
	return data
}
