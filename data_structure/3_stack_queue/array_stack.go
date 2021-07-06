package stack

import (
	"errors"
)

type ArrayStack struct {
	datas []interface{}
	len   int // 栈长度
	limit int // 栈可存储最大长度
}

func NewArrayStack(limit int) *ArrayStack {
	datas := make([]interface{}, 0)

	return &ArrayStack{
		datas,
		0,
		limit,
	}
}

// 存储
// 时间复杂度O(1)
func (stack *ArrayStack) Push(data interface{}) error {
	if stack.len >= stack.limit {
		return errors.New("顺序栈已达到存储上限")
	}

	stack.datas = append(stack.datas, data)
	stack.len++
	return nil
}

// 弹出
// 时间复杂度O(1)
func (stack *ArrayStack) Pop() interface{} {
	if stack.len == 0 {
		return nil
	}

	index := stack.len - 1
	value := stack.datas[index]
	stack.datas = stack.datas[0:index]
	stack.len--

	return value
}

// 翻转
// 时间复杂度O(n)
func (stack *ArrayStack) Reverse() {

	len := stack.len
	datas := make([]interface{}, 0)
	value := stack.Pop()
	for value != nil {
		datas = append(datas, value)
		value = stack.Pop()
	}

	stack.datas = datas
	stack.len = len
}

func (stack *ArrayStack) top() interface{} {
	if stack.len == 0 {
		return nil
	}

	return stack.datas[0]
}
