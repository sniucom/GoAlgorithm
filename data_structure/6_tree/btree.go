package tree

import "fmt"

const M int = 4             // B数阶数
const Min int = (M / 2) - 1 // 节点最少关键词个数

type BTree struct {
	parent *BTree        // 当前节点的父节点
	len    int           // 已存在的关键字个数
	key    [M + 1]int    // 关键词
	child  [M + 1]*BTree // 子树指针
}

// 查找关键字返回索引和节点
func (tree *BTree) Search(data int) (index int, result *BTree) {
	node := &BTree{}
	if tree == nil {
		return -1, node
	}

	var i int
	pointer := tree
	for pointer != nil {
		for i = pointer.len; i > 0 && data <= pointer.key[i]; i-- {
			if data == pointer.key[i] { //
				return i, pointer
			}
		}

		if pointer.child[i] == nil {
			node = pointer
		}

		pointer = pointer.child[i]
	}

	return -1, node
}

// 节点分裂
func (tree *BTree) Split() (result *BTree) {
	node := &BTree{}

	parent := tree.parent
	if parent == nil {
		parent = &BTree{}
	}

	middle := M/2 + 1     // 分裂上移的中间索引
	node.len = M - middle // 右侧新节点数量
	tree.len = middle - 1 // 原节点长度

	// 节点分裂
	for i := 0; i <= node.len; i++ {
		if i < node.len {
			node.key[i+1] = tree.key[middle+i+1]
		}
		node.child[i] = tree.child[middle+i]
	}
	node.parent = parent
	tree.parent = parent

	// 将中间节点插入父节点
	index := parent.len
	for ; tree.key[middle] < parent.key[index]; index-- {
		parent.key[index+1] = parent.key[index]
		parent.child[index+1] = parent.child[index]
	}
	parent.key[index+1] = tree.key[middle]
	parent.child[index] = tree
	parent.child[index+1] = node
	parent.len++

	// 如果父节点超出最大容量 则父节点继续分裂
	if parent.len >= M {
		return parent.Split()
	}

	return parent
}

// 节点调整
func (tree *BTree) Restore() (result *BTree, ok bool) {
	var brother *BTree
	parent := tree.parent
	index := 0
	for ; parent.child[index] != tree; index++ {
		if index > 0 { // 左邻
			brother = parent.child[index-1]
			if brother.len > (M-1)/2 {
				for i := tree.len; i >= 0; i-- {
					tree.key[i+1] = tree.key[i]
				}
				tree.key[1] = parent.key[index]
				parent.key[index] = brother.key[brother.len]
				tree.len++
				brother.len--

				return parent, true
			}

		}

		if index < parent.len {
			brother = parent.child[index+1]
			if brother.len > (M-1)/2 {
				tree.key[tree.len+1] = parent.key[index+1]
				parent.key[index+1] = brother.key[index]
				for i := 1; i < brother.len; i++ {
					brother.key[i] = brother.key[i+1]
				}
				tree.len++
				brother.len--

				return parent, true
			}
		}
	}
	return tree, false
}

// 节点合并
func (tree *BTree) Merge() *BTree {
	var brother *BTree
	parent := tree.parent
	index := 0
	for ; parent.child[index] != tree; index++ {
		if index > 0 { // 与左兄弟合并
			brother = parent.child[index-1]
			brother.len++
			brother.key[brother.len] = parent.key[index]
			for i := 1; i <= tree.len; i++ {
				brother.len++
				brother.key[brother.len] = tree.key[index]
			}

			parent.len--
			for i := index; i < parent.len; i++ {
				parent.key[i] = parent.key[i+1]
				parent.child[i] = parent.child[i+1]
			}
		} else { // 与右兄弟合并
			brother = parent.child[index+1]
			tree.len++
			tree.key[tree.len] = parent.key[index]
			for i := 1; i < tree.len; i++ {
				tree.len++
				tree.key[tree.len] = brother.key[i]
			}
		}

		parent.len--
		for i := index; i <= parent.len; index++ { //改变父节点
			parent.key[index] = parent.key[index+1]
			parent.child[index] = parent.child[index+1]
		}
	}

	return parent
}

func (tree *BTree) FindAfterMinNode(index int) (*int, *BTree) {
	leaf := tree
	if leaf == nil {
		return nil, nil
	} else {
		leaf = leaf.child[index]
		for leaf.child[0] != nil {
			leaf = leaf.child[0]
		}
	}
	return &leaf.key[1], leaf
}

// 遍历
func (tree *BTree) Traverse() {
	queue := []*BTree{}
	queue = append(queue, tree)
	i := 0
	for i < len(queue) {
		current := queue[i]
		i++

		for k := 1; k <= current.len; k++ {
			fmt.Printf(" %d \n", current.key[k])
		}

		for k := 0; k <= current.len; k++ {
			if current.child[k] != nil {
				queue = append(queue, current.child[k])
			}
		}
	}
}

// 插入
func (tree *BTree) Instert(data int) *BTree {
	index, node := tree.Search(data)
	if index < 0 { // 节点不存在
		//  空树
		if tree == nil {
			node.key[1] = data
			node.len = 1
			return node
		}

		var i int
		node.key[0] = data
		for i = node.len; i > 0 && data < node.key[i]; i-- {
			node.key[i+1] = node.key[i]
		}
		node.key[i+1] = data
		node.len++

		if node.len < M {
			return tree
		}

		parent := node.Split()
		for parent.parent != nil {
			parent = parent.parent
		}

		return parent
	}

	return tree
}

func (tree *BTree) Delete(value int) *BTree {
	index, node := tree.Search(value)
	if index != -1 {
		tree = node.DeleteNode(value, index)
	}
	return tree
}

//删除关键字
func (tree *BTree) DeleteNode(value int, index int) *BTree {
	if tree.child[index] != nil { //非终端节点
		valueTemp, nodeTemp := tree.FindAfterMinNode(index)
		tree.key[index] = *valueTemp
		nodeTemp.DeleteNode(*valueTemp, 1)
	} else {
		for i := index; i < tree.len; i++ {
			tree.key[index] = tree.key[index+1]
			tree.child[index] = tree.child[index+1]
		}
		tree.len--
		if tree.len < (M-1)/2 && tree.parent != nil {
			tree, ok := tree.Restore()
			if !ok {
				tree = tree.Merge()
			}
		}
	}

	for tree.parent != nil {
		tree = tree.parent
	}

	return tree
}
