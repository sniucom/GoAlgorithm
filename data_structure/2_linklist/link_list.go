package linklist

import (
	"fmt"
)

// 单向链表
type LinkNode struct {
	Data string
	Next *LinkNode
}

type ILinkNode interface {
	GetLength() int                        // 获取链表长度
	Shift(data string)                     // 向头部添加数据
	Push(data string)                      // 向尾部添加数据
	Insert(index int, data string)         // 向在指定位置插入
	Delete(index int)                      // 按索引删除 并返回被删除的内容
	DeleteNode(data string)                // 按节点删除
	Search(data string) int                // 搜索获取索引位置
	Reverse()                              // 翻转链表
	GetAll() []string                      // 获取所有数据
	OrderUnique()                          // 无序链表顺序去重
	RecurrenceUnique()                     // 无序链表递归去重
	GetString() string                     // 拼接string
	Sort() *LinkNode                       // 排序
	Splice(start int, length int) []string // 截取
	Max() (int, string)                    // 求最大值
}

// 获取链表长度
// 时间复杂度O(n)
func (head *LinkNode) GetLength() int {
	current := head

	length := 0
	for current != nil {
		fmt.Printf("%d", length)
		length++
		current = current.Next
	}

	return length
}

// 向头部添加数据
// 创建一个新节点，将此节点下一节点指向链表第一个节点
// 时间复杂度：O(1)
func (head *LinkNode) Shift(data string) {
	temp := *head
	node := LinkNode{data, &temp}
	*head = node
}

// 向尾部添加数据
// 从链表第一个节点遍历到最后一个节点，向最后一个节点添加到新节点的指针
// 时间复杂度：O(n)
func (head *LinkNode) Push(data string) {
	current := head

	for current.Next != nil {
		current = current.Next
	}

	node := LinkNode{data, nil}
	current.Next = &node
}

// 指定位置插入节点
// 时间复杂度Q(2n)
func (head *LinkNode) Insert(index int, data string) {
	index--
	length := head.GetLength()
	if index < 0 || index > length {
		fmt.Println("插入位置超出索引范围")
		return
	}

	current := head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	node := LinkNode{data, current.Next}
	current.Next = &node
}

// 删除节点
// 时间复杂度Q(2n)
func (head *LinkNode) Delete(index int) string {
	length := head.GetLength()

	if index < 0 || index > length {
		return ""
	}

	current := head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	data := current.Next.Data
	current.Next = current.Next.Next

	return data
}

// 搜索数据并删除
// 时间复杂度O(2n)
func (head *LinkNode) DeleteNode(data string) {
	current := head

	for current.Next != nil && current.Data != data {
		current = current.Next
	}

	current.Next = current.Next.Next
}

// 搜索值获取索引
// 时间复杂度O(n)
func (head *LinkNode) Search(data string) int {
	index := -1
	current := head
	i := 0

	for {
		if current.Data == data {
			index = i
			break
		}

		if current.Next == nil {
			break
		}

		current = current.Next
		i++
	}

	return index
}

// 翻转链表
// 时间复杂度O(n)
func (head *LinkNode) Reverse() {
	if head == nil || head.Next == nil {
		return
	}

	reverseHead := *head
	reverseHead.Next = nil
	current := head.Next
	for current != nil {
		previous := reverseHead
		reverseHead = *current
		reverseHead.Next = &previous

		current = current.Next
	}

	*head = reverseHead
}

// 获取所有数据转为slice
// 时间复杂度O(n)
func (head *LinkNode) GetAll() []string {
	list := make([]string, 0, head.GetLength())

	current := head
	for current.Next != nil {
		list = append(list, current.Data)
	}
	list = append(list, current.Data)

	return list
}

// 无序链表顺序删除重复
// 时间复杂度O(n2)
func (head *LinkNode) OrderUnique() {
	if head == nil || head.Next == nil {
		return
	}

	outer := head          // 外层循环
	var inner *LinkNode    // 内层循环
	var previous *LinkNode // 前驱节点

	for outer != nil && outer.Next != nil {
		inner = outer.Next
		previous = outer
		for {
			if outer.Data == inner.Data {
				previous.Next = inner.Next
			} else {
				previous = inner
			}

			if inner.Next == nil {
				break
			}
			inner = inner.Next
		}

		outer = outer.Next
	}
}

// 无序链表递归去重
// 时间复杂度O(n2)
func (head *LinkNode) RecurrenceUnique() {
	if head == nil || head.Next == nil {
		return
	}

	values := map[string]int{head.Data: 1}
	head.Next = RecurrenceNext(values, head.Next)
}

// 递归获取下一节点
func RecurrenceNext(values map[string]int, node *LinkNode) *LinkNode {
	if node == nil {
		return node
	}

	_, ok := values[node.Data]
	if !ok {
		values[node.Data] = 1
		node.Next = RecurrenceNext(values, node.Next)
		return node
	}

	return RecurrenceNext(values, node.Next)
}

// 获取拼接字符串
func (head *LinkNode) GetString() string {
	result := ""
	current := head
	for {
		result += current.Data
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	return result
}

// 链表排序
// 时间复杂度O(n2)
func (head *LinkNode) Sort() *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := &LinkNode{"_start_", nil}
	outer := head
	var inner *LinkNode
	var next *LinkNode
	for outer != nil {
		inner = node
		next = outer.Next
		for inner != nil {
			if inner.Next == nil {
				outer.Next = nil
				inner.Next = outer
				break
			}

			if (inner.Data == "_start_" || outer.Data > inner.Data) && outer.Data < inner.Next.Data {
				outer.Next = inner.Next
				inner.Next = outer
				break
			}
			inner = inner.Next
		}
		outer = next
	}

	return node.Next
}

// 截取数据
// 时间复杂度O(2n)
func (head *LinkNode) Splice(start int, length int) []string {
	listLength := head.GetLength()
	if start < 0 || start > (listLength-1) {
		fmt.Println("截取起始超出")
		return nil
	}

	index := 0
	end := start + length
	data := []string{}
	current := head
	for current.Next != nil && index < listLength && index < end {
		if index >= start {
			data = append(data, current.Data)
		}

		index++
		current = current.Next
	}

	return data
}

// 获取最大节点
// 时间复杂度O(n)
func (head *LinkNode) Max() (int, string) {
	if head == nil {
		return -1, ""
	}

	if head.Next == nil {
		return 0, head.Data
	}

	index := 0
	first := -1
	maxData := head.Data
	current := head.Next
	for current != nil {
		if current.Data > maxData {
			first = index
			maxData = current.Data
		}
		index++
		current = current.Next
	}

	return first, maxData
}
