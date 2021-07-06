package stack

import (
	"errors"
)

type MaxMinStackNode struct {
	data int // 数据
	min  int // 最小值
	max  int // 最大值
}

type MaxMinStack struct {
	datas []*MaxMinStackNode
	len   int // 栈长度
	limit int // 栈可存储最大长度
}

func NewMaxMinStack(limit int) *MaxMinStack {

	datas := make([]*MaxMinStackNode, 0)

	return &MaxMinStack{
		datas,
		0,
		limit,
	}
}

// 存储
// 时间复杂度O(1)
func (stack *MaxMinStack) Push(data int) error {
	len := stack.len

	if len >= stack.limit {
		return errors.New("栈已达到存储上限")
	}

	min, max := data, data
	if len > 0 {
		top := stack.datas[len-1]
		topMin := top.min
		topMax := top.max
		if min > topMin {
			min = topMin
		}
		if max < topMax {
			max = topMax
		}
	}

	node := &MaxMinStackNode{data, min, max}

	stack.datas = append(stack.datas, node)
	stack.len++
	return nil
}

// 弹出
// 时间复杂度O(1)
func (stack *MaxMinStack) Pop() (data int, ok bool) {
	if stack.len == 0 {
		return
	}

	index := stack.len - 1
	node := stack.datas[index]
	stack.datas = stack.datas[0:index]

	stack.len--

	return node.data, true
}

// 获取最小值
// 时间复杂度O(1)
func (stack *MaxMinStack) Min() (min int, ok bool) {
	if stack.len == 0 {
		return
	}

	return stack.datas[stack.len-1].min, true
}

// 获取最大值
// 时间复杂度O(1)
func (stack *MaxMinStack) Max() (max int, ok bool) {
	if stack.len == 0 {
		return
	}

	return stack.datas[stack.len-1].max, true
}

// 获取顶部数据
// 时间复杂度O(1)
func (stack *MaxMinStack) Top() (max int, ok bool) {
	if stack.len == 0 {
		return
	}

	return stack.datas[stack.len-1].data, true
}
