package tree

import (
	"errors"
	"fmt"
)

const BPM = 5                 // 4阶B+树
const BPLIMIT = (BPM + 1) / 2 // 分裂点
const BPMIN = M/2 - 1         // 最小存储

// B+树
type BPlusTree struct {
	root *BPlusTreeNode
}

// 节点
type BPlusTreeNode struct {
	IsLeaf bool                    // 是否是叶子节点
	Len    int                     // 节点当前存储长度
	Key    [BPM + 1]int            // 关键字数组
	Items  [BPM + 1]*BPlusTreeNode // 孩子节点指针
	Leaf   *BPlusTreeLeaf          // 叶子节点存储
	Parent *BPlusTreeNode
}

// 叶子节点链表结构
type BPlusTreeLeaf struct {
	Prev  *BPlusTreeNode       // 前一节点
	Next  *BPlusTreeNode       // 后一节点
	Value [BPM + 1]interface{} // 存储数据
}

// 插入
// 1. 空树 创建叶子节点并作为
// 2. 非空树，查找插入节点，插入，如果超过M阶，则分裂并进位
func (tree *BPlusTree) Insert(id int, data interface{}) {
	leaf := &BPlusTreeLeaf{}
	node := &BPlusTreeNode{}

	// 空树 直接创建叶子节点作为根
	if tree.root == nil {
		leaf.Value[0] = data
		node.IsLeaf = true
		node.Len++
		node.Key[1] = id
		node.Leaf = leaf
		tree.root = node
		return
	}

	insertIndex, insertNode := tree.Search(id)
	fmt.Println("\n----------------")
	fmt.Printf("可插入节点 %d %v \n", insertIndex, insertNode)
	if insertIndex > -1 || !insertNode.IsLeaf { // 索引已存在 或 非叶子节点
		return
	}

	// 插入数据
	var i int
	insertNode.Key[0] = id
	for i = insertNode.Len; i > 0 && id < insertNode.Key[i]; i-- {
		// 索引0是留空的， len+1的索引在插入是肯定是nil， 因为如果len+ != nil 代表节点满了， 要分裂了
		insertNode.Key[i+1] = insertNode.Key[i] // 向后移
	}
	insertNode.Key[i+1] = id
	insertNode.Leaf.Value[i] = data
	insertNode.Len++

	if insertNode.Len >= BPM { // 节点存储宽度超出限制 则分裂
		parent := insertNode.Split()
		for parent.Parent != nil {
			parent = parent.Parent
		}
		tree.root = parent
	}
}

// 节点分裂
func (splitNode *BPlusTreeNode) Split() *BPlusTreeNode {
	rightNode := &BPlusTreeNode{} // 新分裂出来的右节点
	if splitNode.IsLeaf {
		rightNodeLeaf := &BPlusTreeLeaf{}
		rightNode.IsLeaf = true
		rightNode.Leaf = rightNodeLeaf
	}

	parent := splitNode.Parent // 父节点
	if parent == nil {
		parent = &BPlusTreeNode{}
	}

	middle := BPLIMIT // 中间节点索引起始  5阶左边保留0 1 2  右边保留 3 4
	// 填充右节点
	index := 0
	for i := middle; i <= splitNode.Len; i++ {
		rightNode.Key[index+1] = splitNode.Key[i+1] // 索引分裂 123,45
		if splitNode.IsLeaf {                       // 叶子节点
			rightNode.Leaf.Value[index] = splitNode.Leaf.Value[i] // 值跟列 012， 34
		} else { // 非叶子节点
			rightNode.Items[index] = splitNode.Items[i] // 孩子节点指针分裂
		}
		splitNode.Len--
		rightNode.Len++
		index++
	}
	rightNode.Parent = parent
	splitNode.Parent = parent
	fmt.Printf("分裂左 %v \n 分裂右 %v Middle: %d", splitNode, rightNode, middle)

	// 如果是叶子节点，需要转移指针
	if splitNode.IsLeaf {
		next := splitNode.Leaf.Next
		splitNode.Leaf.Next = rightNode
		rightNode.Leaf.Prev = splitNode
		rightNode.Leaf.Next = next
	}

	// 将该节点插入到父节点
	j := parent.Len
	for ; j > 0 && splitNode.Key[middle+1] < parent.Key[j]; j-- {
		parent.Key[j+1] = parent.Key[j]
		parent.Items[j] = parent.Items[j-1]
	}
	parent.Key[j+1] = splitNode.Key[middle+1]
	parent.Items[j] = splitNode
	parent.Items[j+1] = rightNode
	parent.Len++

	fmt.Printf("父节点 %v \n", parent)

	if parent.Len >= BPM {
		parent = parent.Split()
	}

	return parent
}

// 搜索可插入节点
func (tree *BPlusTree) Search(id int) (index int, result *BPlusTreeNode) {
	if tree == nil {
		return -1, nil
	}

	node := &BPlusTreeNode{}
	var i int
	pointer := tree.root
	for pointer != nil {
		for i = pointer.Len; i > 0 && id <= pointer.Key[i]; i-- {
			if pointer.IsLeaf && id == pointer.Key[i] { // 在叶子节点中找到所能因直接返回
				return i, pointer
			}
		}

		if pointer.IsLeaf && pointer.Leaf.Value[i] == nil {
			node = pointer
		}
		fmt.Printf("子节点 %v  \n", pointer)
		pointer = pointer.Items[i]
	}

	return -1, node
}

// 更新
func (tree *BPlusTree) Update(id int, data interface{}) (result interface{}, err error) {
	index, node := tree.Search(id)
	if index < 0 {
		return nil, errors.New("您要更新的数据不存在")
	}

	node.Leaf.Value[index] = data

	return node, nil
}

// 简单范围查找
func (tree *BPlusTree) Query(compare string, compareId int, limit int) []interface{} {

	index := -1
	pointer := tree.root
	if compare != "!=" {
		index, pointer = tree.Search(compareId)
	}

	list := make([]interface{}, 0)
	var start *BPlusTreeLeaf
	var count int
	if pointer != nil && pointer.IsLeaf {
		switch compare {
		case "=":
			if index >= 0 {
				list = append(list, pointer.Leaf.Value[index])
			}
		case ">", ">=":
			if compare == ">" {
				pointer = pointer.Leaf.Next
			}

			count = 0
			for start != nil {
				if limit > 0 && count >= limit {
					break
				}

				for value := range pointer.Leaf.Value {
					if limit > 0 && count >= limit {
						break
					}
					list = append(list, value)
					count++
				}
				pointer = pointer.Leaf.Next
			}
		case "<", "<=":
			if compare == "<" {
				pointer = pointer.Leaf.Prev
			}
			count = 0
			for start != nil {
				if limit > 0 && count >= limit {
					break
				}

				for value := range pointer.Leaf.Value {
					if limit > 0 && count >= limit {
						break
					}
					list = append(list, value)
					count++
				}
				pointer = pointer.Leaf.Next
			}
		}
	}

	return list
}

// 测试数据
type Book struct {
	ID       int     // 唯一标识
	Title    string  // 书名
	Price    float32 // 价格
	Describe string  // 描述
}

type Librarian struct {
	SN int
}

func (librarian *Librarian) Record(title string, price float32, Describe string) *Book {
	librarian.SN++
	return &Book{
		librarian.SN,
		title,
		price,
		Describe,
	}
}
