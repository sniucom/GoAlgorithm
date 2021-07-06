package stack

import (
	"container/list"
	"errors"
)

type LinkedListStack struct {
	datas *list.List // 存储的链表
	len   int        // 栈长度
	limit int        // 栈最大可存储长度
}

func NewLinkedListStack(limit int) *LinkedListStack {
	return &LinkedListStack{
		list.New(),
		0,
		limit,
	}
}

// 存储
// 时间复杂度O(1)
func (stack *LinkedListStack) Push(data interface{}) error {
	if stack.len >= stack.limit {
		return errors.New("顺序栈已达到存储上限")
	}

	stack.datas.PushFront(data)
	stack.len++

	return nil
}

// 弹出
// 时间复杂度O(1)
func (stack *LinkedListStack) Pop() interface{} {
	if stack.len == 0 {
		return nil
	}

	stack.len--
	value := stack.datas.Front().Value
	stack.datas.Remove(stack.datas.Front())

	return value
}

// 栈翻转
// 时间复杂度O(n)
func (stack *LinkedListStack) Reverse() {
	if stack.len == 0 {
		return
	}

	current := stack.datas.Back().Prev()
	var prev *list.Element
	for current != nil {
		prev = current.Prev()
		stack.datas.MoveToBack(current)
		current = prev
	}
}
