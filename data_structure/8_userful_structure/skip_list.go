package useful_structure

import (
	"fmt"
	"math/rand"
)

const MAX_LEVEL = 32

type SkipListNode struct {
	Key   int
	Value interface{}
	next  []*SkipListNode
}

type SkipList struct {
	limit int           // 层级限制
	max   int           // 最大层级
	len   int           // 跳表存储长度
	head  *SkipListNode // 头元素
}

func NewSkipList() *SkipList {
	return &SkipList{
		limit: MAX_LEVEL,
	}
}

// 随机插入层级
func (list *SkipList) random() int {
	level := 1
	probability := list.limit / 4
	for level < list.limit && rand.Intn(list.limit) < probability {
		level++
	}

	return level
}

// 搜索转折节点
func (list *SkipList) Search(key int) []*SkipListNode {
	pointer := list.head
	prevs := make([]*SkipListNode, MAX_LEVEL+1)

	count := 0
	for i := list.max - 1; i >= 0; i-- {
		for pointer.next[i] != nil && pointer.next[i].Key < key {
			pointer = pointer.next[i]
			count++
		}

		prevs[i] = pointer
		count++
	}

	fmt.Printf("寻找%d查询%d次---\n", key, count)

	return prevs
}

// 查找
func (list *SkipList) Find(key int) *SkipListNode {
	prevs := list.Search(key)
	if prevs[0] != nil && prevs[0].next[0] != nil && prevs[0].next[0].Key == key {
		return prevs[0].next[0]
	}

	return nil
}

// 插入
func (list *SkipList) Insert(key int, data interface{}) bool {
	prevs := list.Search(key)
	if prevs[0] != nil && prevs[0].next[0] != nil && prevs[0].next[0].Key == key {
		return false
	}

	pointer := &SkipListNode{
		key,
		data,
		make([]*SkipListNode, MAX_LEVEL+1),
	}

	if prevs[0] == nil {
		list.head = pointer
		list.max = 1
		list.len = 1
		return true
	}

	// 为该节点随机产生层数
	level := list.random()
	if level > list.max {
		for i := list.max; i < level; i++ {
			prevs[i] = list.head
		}
		list.max = level
	}

	for i := level - 1; i >= 0; i-- {
		pointer.next[i] = prevs[i].next[i]
		prevs[i].next[i] = pointer
	}

	list.len++

	return true
}

// 更新
func (list *SkipList) Update(key int, data interface{}) bool {
	node := list.Find(key)
	if node == nil {
		return false
	}

	node.Value = data

	return true
}

// 删除
func (list *SkipList) Delete(key int) bool {
	prevs := list.Search(key)
	if prevs[0] == nil || prevs[0].Key != key { // 找不到
		return false
	}

	node := prevs[0].next[0]

	for i := 0; i < list.max; i++ {
		if prevs[i].next[i] == node {
			prevs[i].next[i] = node.next[i]
		}
	}

	for list.max > 1 && list.head.next[list.max-1] == nil {
		list.max--
	}

	return true
}

func (list *SkipList) Print() {
	pointer := list.head
	keys := make([]int, 0)
	for pointer != nil {
		keys = append(keys, pointer.Key)
		pointer = pointer.next[0]
	}

	fmt.Printf("%v\n", keys)
}
