package tree

import "fmt"

type ArrayTree struct {
	items []interface{}
	Len   int
	Mark  int
}

func NewArrayTree(root int) *ArrayTree {
	items := make([]interface{}, 2048)
	items[0] = root
	return &ArrayTree{
		items,
		1,
		1,
	}
}

// 添加
func (tree *ArrayTree) Add(data int) {
	index := tree.InsertIndex(data, 0)
	tree.items[index] = data
	tree.Len++
	if index > tree.Mark {
		tree.Mark = index + 2
	}
	fmt.Printf(" 添加%d\n 当前存储:%v\n\n", data, tree.items[:tree.Mark])
}

// 值是否存在
func (tree *ArrayTree) Exists(data int) bool {
	index := tree.GetIndex(data, 0)
	if index >= 0 {
		return true
	} else {
		return false
	}
}

// 移除元素
func (tree *ArrayTree) Remove(data int) {
	index := tree.GetIndex(data, 0)
	tree.RemoveChild(index)
	fmt.Printf("移除元素%d后 当前存储%v \n", data, tree.items[:tree.Mark])
}

func (tree *ArrayTree) RemoveChild(index int) {
	if index+2 <= tree.Mark {
		if tree.items[2*index+1] != nil {
			tree.Len--
			tree.RemoveChild(2*index + 1)
		}
		if tree.items[2*index+2] != nil {
			tree.Len--
			tree.RemoveChild(2*index + 2)
		}
	}

	tree.items[index] = nil
}

func (tree *ArrayTree) GetIndex(data int, parent int) int {
	index := -1
	if tree.items[parent] != nil {
		parentValue := tree.items[parent].(int)
		if data > parentValue {
			index = tree.GetIndex(data, 2*parent+2)
		} else if data < parentValue {
			index = tree.GetIndex(data, 2*parent+1)
		} else {
			index = parent
		}
	}

	return index
}

// 获取插入索引
func (tree *ArrayTree) InsertIndex(data int, parent int) int {
	index := -1
	parentValue := tree.items[parent].(int)

	if data > parentValue {
		index = 2*parent + 2
	} else if data < parentValue {
		index = 2*parent + 1
	}

	if index > 0 && tree.items[index] != nil {
		index = tree.InsertIndex(data, index)
	}

	return index
}
